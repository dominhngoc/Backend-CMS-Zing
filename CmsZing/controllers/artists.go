package controllers

import (
	"backend-cms-zing/CmsZing/Validation"
	"backend-cms-zing/CmsZing/conf"
	"backend-cms-zing/CmsZing/models"
	"backend-cms-zing/CmsZing/response"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

// ArtistsController operations for Artists
type ArtistsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ArtistsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Artists
// @Param	body		body 	models.ArtistsSwagger	true		"body for Artists content"
// @Success 201 {object} models.ArtistsSwagger
// @Failure 403 221 parse json fail <br> 303 save fail <br> 217 field required <br> 218 min character number required  6 <br> 219 max characters number required 50
// @router / [post]
func (c *ArtistsController) Post() {
	var v models.Artists
	var detailErrorCode = make(map[string]int)
	var responseData = response.Response{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		detailErrorCode["Body"] = conf.PARSE_JSON_FAIL
		responseData = response.DataResponse("403", "parse json fail", nil, detailErrorCode, "False")
		c.Data["json"] = responseData
		c.ServeJSON()
		return
	}

	//validate form
	if isPass,detailErrorCode := Validation.RebuildValidate(v); isPass==false {
		responseData = response.DataResponse("403", "Validate False", nil, detailErrorCode, "False")
		c.Data["json"] = responseData
		c.ServeJSON()
		return
	}

	// save file to storage/image from storage/temp

	// add song database
	id,err := models.AddArtists(&v)
	if  err != nil {
		detailErrorCode["System"] = conf.SAVE_FAIL
		responseData = response.DataResponse("403", err.Error(), nil, detailErrorCode, "False")
		c.Data["json"] = responseData
		c.ServeJSON()
		return
	}
	v.Id=int(id)
	responseData = response.DataResponse("200", "Add success" , v, nil, "true")
	c.Data["json"] = responseData
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Artists by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} response.ArtistSwaggerResponse <br>Return success
// @Failure 403 850 something went wrong <br> 305 not exists <br> 809 must be integer
// @router /:id [get]
func (c *ArtistsController) GetOne() {
	responseData := response.Response{}
	DetailError := make(map[string]int)

	//get id paramater
	idStr := c.Ctx.Input.Param(":id")

	// check id is integer
	id, ParseErr := strconv.Atoi(idStr)
	if ParseErr != nil {
		DetailError["id"] = conf.INCORRECT_FORMAT
		responseData = response.DataResponse("403", "get artist fail", nil, DetailError, "false")
		c.Data["json"] = responseData
		c.ServeJSON()
		return
	}

	// get data from database
	data, err := models.GetArtistsById(id)

	if err != nil {
		DetailError["system"] = conf.SOMETHING_WRONG
		responseData = response.DataResponse("403", "get artist fail", nil, DetailError, "false")
	}

	// empty data
	if data == nil {
		DetailError["artist"] = conf.NOT_FOUND
		responseData = response.DataResponse("403", "get artist fail", nil, DetailError, "false")
	}

	// get data success
	if data != nil {
		DetailError = nil
		responseData = response.DataResponse("200", "return success", data, nil, "true")
	}
	c.Data["json"] = responseData
	c.ServeJSON()
}

// check invalid fields
func ValidateFields(fields []string) (responseData response.Response) {
	o := orm.NewOrm()
	DetailError := make(map[string]int)
	for _, m := range fields {
		sql := "SELECT " + m + " FROM artists"
		if _, err := o.Raw(sql).Exec(); err != nil {
			DetailError["column_name"] = conf.INCORRECT_FIELD
			responseData = response.DataResponse("403", "Wrong field/Column-name "+m, nil, DetailError, "false")
			return responseData
		}
	}
	return responseData
}

// GetAll ...
// @Title Get All
// @Description get Artists
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	searchByName	query	string	false	"keyword"
// @Success 200 {object} response.ArtistSwaggerResponse <br>Return success
// @Failure 403 850 something went wrong <br> 305  not found data <br> 851 incorrect field
// @router / [get]
func (c *ArtistsController) GetAll() {
	var fields []string
	responseData := response.Response{}
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// if data return doesn't have error
	if responseData = ValidateFields(fields); responseData.DetailErrorCode == nil {
		DetailError := make(map[string]int)
		searchByName := c.GetString("searchByName")
		// get artists information
		list, err := models.GetAllArtists(fields, searchByName)
		if err != nil {
			DetailError["system"] = conf.SOMETHING_WRONG
			responseData = response.DataResponse("403", "Something went wrong", nil, DetailError, "false")
		}
		if list == nil {
			DetailError["artist"] = conf.NOT_FOUND
			responseData = response.DataResponse("403", "not found data", nil, DetailError, "false")
		}
		if list != nil {
			DetailError = nil
			responseData = response.DataResponse("200", "return success", list, nil, "true")
		}
	}
	c.Data["json"] = responseData
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Artists
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Artists	true		"body for Artists content"
// @Success 200 {object} models.Artists
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ArtistsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Artists{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateArtistsById(&v); err == nil {
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
// @Description delete the Artists
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 850 something went wrong <br> 305 not exists <br> 809 must be integer <br> id not exists
// @router /:id [delete]
func (c *ArtistsController) Delete() {
	responseData := response.Response{}
	DetailError := make(map[string]int)

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	// check id is integer
	id, ParseErr := strconv.Atoi(idStr)
	if ParseErr != nil {
		DetailError["id"] = conf.INCORRECT_FORMAT
		responseData = response.DataResponse("403", "delete fail", nil, DetailError, "false")
		c.Data["json"] = responseData
		c.ServeJSON()
		return
	}
	// check id exists in database
	data, _ := models.GetArtistsById(id)
	if data == nil {
		DetailError["id"] = conf.NOT_FOUND
		responseData = response.DataResponse("403", "delete fail", nil, DetailError, "false")
		c.Data["json"] = responseData
		c.ServeJSON()
		return
	}
	// delete from artist table and artist_song table
	err1 := models.DeleteTransaction(id)
	if err1 != nil {
		DetailError["system"] = conf.SOMETHING_WRONG
		responseData = response.DataResponse("403", "delete fail", nil, DetailError, "false")
		c.Data["json"] = responseData
		c.ServeJSON()
		return
	}
	c.Data["json"] = "delete success !"
	c.ServeJSON()
	return

}
