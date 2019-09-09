package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	//"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ArtistSong struct {
	Id        int       `orm:"column(id);auto"`
	ArtistId  int       `orm:"column(artist_id)"`
	SongId    int       `orm:"column(song_id)"`
	IsDeleted int       `orm:"column(is_deleted)"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);auto_now"`
}

func (t *ArtistSong) TableName() string {
	return "artist_song"
}

func init() {
	orm.RegisterModel(new(ArtistSong))
}

// AddArtistSong insert a new ArtistSong into database and returns
// last inserted Id on success.
// goi tu model songs de luu vao
func AddArtistSong(id int64, m *SongInfo) (err error) {
	// query in artist_song table
	o := orm.NewOrm()
	index := []interface{}{}
	sqlStr := "INSERT INTO artist_song (artist_id,song_id) VALUES "
	for _, values := range m.Singer {
		sqlStr += "(?,?),"
		index = append(index, values.Id, id)
	}
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	_, err = o.Raw(sqlStr, index).Exec()
	return
}

// GetArtistSongById retrieves ArtistSong by Id. Returns error if
// Id doesn't exist
func GetArtistSongById(id int) (v *ArtistSong, err error) {
	o := orm.NewOrm()
	v = &ArtistSong{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllArtistSong retrieves all ArtistSong matches certain condition. Returns empty list if
// no records exist
func GetAllArtistSong(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ArtistSong))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ArtistSong
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateArtistSong updates ArtistSong by Id and returns error if
// the record to be updated doesn't exist
func UpdateArtistSongById(m *ArtistSong) (err error) {
	o := orm.NewOrm()
	v := ArtistSong{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//DeleteArtistSong deletes ArtistSong by Id and returns error if
//the record to be deleted doesn't exist
func DeleteArtistSong(id int,o orm.Ormer) (err error) {
	sql := "DELETE FROM artist_song WHERE artist_song.artist_id = " + strconv.Itoa(id)
	_, err = o.Raw(sql).Exec()
	return err
}
