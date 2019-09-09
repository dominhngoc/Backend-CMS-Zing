package controllers

import (
	"backend-cms-zing/CmsZing/Validation"
	"backend-cms-zing/CmsZing/conf"
	"backend-cms-zing/CmsZing/models"
	"backend-cms-zing/CmsZing/response"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	uuid2 "github.com/nu7hatch/gouuid"
	"os/exec"
	"strconv"
	"strings"
)

// SongsController operations for Songs
type SongsController struct {
	beego.Controller
}

// URLMapping ...
func (c *SongsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Songs
// @Param	body		body 	models.SongInfo	true		"body for Songs content"
// @Success 201 {object} response.Response add success
// @Failure 403 221 parse json fail <br> 303 save failures <br> 217 field required <br> 218 not enough characters required <br> 219 name max 125 characters
// @router /create [post]
func (c *SongsController) Post() {
	var v models.SongInfo
	var detailErrorCode = make(map[string]int)
	var songs = response.Response{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		detailErrorCode["Body"] = conf.PARSE_JSON_FAIL
		songs = response.DataResponse("403", err.Error(), nil, detailErrorCode, "False")
		c.Data["json"] = songs
		c.ServeJSON()
		return
	}

	//validate form
	if status, detailErrorCode := Validation.RebuildValidate(v); !status {
		songs = response.DataResponse("403", "Validate False", nil, detailErrorCode, "False")
		c.Data["json"] = songs
		c.ServeJSON()
		return
	}

	// move file from storage/temp to storage/image
	if err := MoveFileImage(v.ImageUrl); err != nil {
		detailErrorCode["System"] = conf.SAVE_FAIL
		songs = response.DataResponse("403", err.Error(), nil, detailErrorCode, "False")
	}

	// move file from storage/temp to storage/music
	if err := MoveFileMusic(v.MusicUrl); err != nil {
		detailErrorCode["System"] = conf.SAVE_FAIL
		songs = response.DataResponse("403", err.Error(), nil, detailErrorCode, "False")
	}

	// update image url in DB
	v.ImageUrl = strings.Replace(v.ImageUrl, "temp", "image", 1)
	v.MusicUrl = strings.Replace(v.MusicUrl, "temp", "music", 1)

	if _, err := models.AddSongs(&v); err != nil {
		detailErrorCode["System"] = conf.SAVE_FAIL
		songs = response.DataResponse("403", err.Error(), nil, detailErrorCode, "False")
		c.Data["json"] = songs
		c.ServeJSON()
		return
	}

	songs = response.DataResponse("201", "add succes", nil, nil, "True")
	c.Data["json"] = songs
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Songs by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Songs
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SongsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSongsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Songs
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Songs
// @Failure 403
// @router / [get]
func (c *SongsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllSongs(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Songs
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Songs	true		"body for Songs content"
// @Success 200 {object} models.Songs
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SongsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Songs{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSongsById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Songs
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SongsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSongs(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

/**
** Validate Image
 */
var images = []string{"image/png", "image/jpg"}

func ValidateImage(headerContentType string) bool {
	for _, imageType := range images {
		if imageType == headerContentType {
			return true
		}
	}
	return false
}

// uploadfile
// @Param	image		formData 	file	true		"the image to upload :jpg,png"
// @Success 200 {object} response.Response upload success!
// @Failure 403 800 failures upload <br> 809 Incorrect file format <br> 999 max size is 2MB <br> 305 file not found
// @router /file/upload/image [post]
func (c *SongsController) UploadFileImage() {
	var songs = response.Response{}
	var detailErrorCode = make(map[string]int)

	file, header, _ := c.GetFile("image") // tieu de
	if file == nil {
		detailErrorCode["File"] = conf.NOT_FOUND
		songs = response.DataResponse("403", "failures upload", nil, detailErrorCode, "False")
		c.Data["json"] = songs
		c.ServeJSON()
		return
	}
	if !ValidateImage(header.Header.Get("Content-Type")) {
		detailErrorCode["File"] = conf.INCORRECT_FORMAT
		songs = response.DataResponse("403", "file has .jpg or .png", nil, detailErrorCode, "False")
		c.Data["json"] = songs
		c.ServeJSON()
		return
	}
	// size < 2MB
	if header.Size > 2097152 {
		detailErrorCode["File"] = conf.MAX_SIZE_FILE
		songs = response.DataResponse("403", "file must be less than 2MB", nil, detailErrorCode, "False")
		c.Data["json"] = songs
		c.ServeJSON()
		return
	}

	uuid, _ := uuid2.NewV4()
	rand := uuid.String()
	filename := header.Filename
	if header.Header.Get("Filename") != "" {
		filename = header.Header.Get("Filename")
	}
	path := "/storage/temp/" + rand + "_" + filename
	err := c.SaveToFile("image", "."+path)

	imageUrl := conf.BaseServer + path

	if err != nil {
		detailErrorCode["File"] = conf.FAIL_UPLOAD
		songs = response.DataResponse("403", err.Error(), nil, detailErrorCode, "False")
		c.Data["json"] = songs
		c.ServeJSON()
		return
	}

	songs = response.DataResponse("200", "upload file success", imageUrl, nil, "True")
	c.Data["json"] = songs
	c.ServeJSON()
}

// uploadfile
// @Param	music		formData 	file	true		"the image to upload : .mp3"
// @Success 200 {object} response.Response upload success!
// @Failure 403 800 failures upload <br> 809 Incorrect file format <br> 305 file not found
// @router /file/upload/music [post]
func (c *SongsController) UploadFileMusic() {
	var songs = response.Response{}
	var detailErrorCode = make(map[string]int)
	file, header, _ := c.GetFile("music") // content
	if file == nil {
		detailErrorCode["File"] = conf.NOT_FOUND
		songs = response.DataResponse("403", "failures upload", nil, detailErrorCode, "False")
		c.Data["json"] = songs
		c.ServeJSON()
		return

	}
	if header.Header.Get("Content-Type") != "audio/mp3" {
		detailErrorCode["File"] = conf.INCORRECT_FORMAT
		songs = response.DataResponse("403", "file has .mp3", nil, detailErrorCode, "False")
		c.Data["json"] = songs
		c.ServeJSON()
		return
	}

	uuid, _ := uuid2.NewV4()
	rand := uuid.String()
	filename := header.Filename
	if header.Header.Get("Filename") != "" {
		filename = header.Header.Get("Filename")
	}
	path := "/storage/temp/" + rand + "_" + filename
	err := c.SaveToFile("music", "."+path)

	imageUrl := conf.BaseServer + path

	if err != nil {
		detailErrorCode["System"] = conf.FAIL_UPLOAD
		songs = response.DataResponse("403", err.Error(), nil, detailErrorCode, "False")
		c.Data["json"] = songs
		c.ServeJSON()
		return
	}

	songs = response.DataResponse("200", "upload file success", imageUrl, nil, "True")
	c.Data["json"] = songs
	c.ServeJSON()
}

func MoveFileImage(imageUrl string) (err error) {
	imageName := strings.Split(imageUrl, "http://27.72.88.246:8280/storage/temp/")
	fileUrl := "storage/temp/" + imageName[1]
	command := exec.Command("mv", fileUrl, "storage/image")
	err = command.Run()
	// bat loi
	if err != nil {
		return err
	}

	return nil
}

func MoveFileMusic(imageUrl string) (err error) {
	imageName := strings.Split(imageUrl, "http://27.72.88.246:8280/storage/temp/")
	fileUrl := "storage/temp/" + imageName[1]
	command := exec.Command("mv", fileUrl, "storage/music")
	err = command.Run()
	//bat loi
	if err != nil {
		return err
	}

	return nil
}
