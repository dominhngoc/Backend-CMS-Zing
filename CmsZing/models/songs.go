package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Songs struct {
	Id            int       `orm:"column(id);auto"`
	Name          string    `orm:"column(name);size(255)"`
	Lyrics        string    `orm:"column(lyrics);null"`
	MusicUrl      string    `orm:"column(music_url);size(255)"`
	Kind          string    `orm:"column(kind);size(255)"`
	ImageUrl      string    `orm:"column(image_url);size(255);null"`
	LrcUrl        string    `orm:"column(lrc_url);size(255);null"`
	Released      string    `orm:"column(released);size(255);null"`
	CreatedUserId int       `orm:"column(created_user_id)"`
	UpdatedUserId int       `orm:"column(updated_user_id)"`
	IsDeleted     int       `orm:"column(is_deleted)"`
	CreatedAt     time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt     time.Time `orm:"column(updated_at);type(timestamp);auto_now"`
	SingerName    string    `orm:"column(artistName);size(255)"`
	IdSinger      int       `orm:"column(artistId)"`
}

// get info songs
type SongInfo struct {
	Id       int      `orm:"column(id);auto"`
	Name     string   `orm:"column(name);size(255)" valid:"Required;MinSize(6);MaxSize(125)"`
	Lyrics   string   `orm:"column(lyrics);null" valid:"Required;MinSize(100)"`
	MusicUrl string   `orm:"column(music_url);size(255)" valid:"Required"`
	Kind     string   `orm:"column(kind);size(255)" valid:"Required"`
	ImageUrl string   `orm:"column(image_url);size(255);null"`
	LrcUrl   string   `orm:"column(lrc_url);size(255);null"`
	Released string   `orm:"column(released);size(255);null"`
	Singer   []Singer `valid:"Required"`
}

// get singer
type Singer struct {
	Id   int
	Name string
}

func (t *Songs) TableName() string {
	return "songs"
}

func init() {
	orm.RegisterModel(new(Songs))
}

// AddSongs insert a new Songs into database and returns
// last inserted Id on success.
func AddSongs(m *SongInfo) (id int64, err error) {
	o := orm.NewOrm()

	// query in song table
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.InsertInto("songs", "songs.name", "songs.lyrics", "songs.music_url", "songs.kind", "songs.image_url", "songs.lrc_url", "songs.released").
		Values("?", "?", "?", "?", "?", "?", "?")
	sql := qb.String()

	// transaction
	err = o.Begin()
	i, error := o.Raw(sql, m.Name, m.Lyrics, m.MusicUrl, m.Kind, m.ImageUrl, m.LrcUrl, m.Released).Exec()

	if error != nil {
		err = o.Rollback()
	}

	id, _ = i.LastInsertId()
	if errs := AddArtistSong(id, m); errs != nil {
		err = o.Rollback()
	}

	err = o.Commit()

	return
}

// GetSongsById retrieves Songs by Id. Returns error if
// Id doesn't exist
func GetSongsById(id int) (v *Songs, err error) {
	o := orm.NewOrm()
	v = &Songs{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSongs retrieves all Songs matches certain condition. Returns empty list if
// no records exist
func GetAllSongs(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Songs))
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

	var l []Songs
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

// UpdateSongs updates Songs by Id and returns error if
// the record to be updated doesn't exist
func UpdateSongsById(m *Songs) (err error) {
	o := orm.NewOrm()
	v := Songs{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSongs deletes Songs by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSongs(id int) (err error) {
	o := orm.NewOrm()
	v := Songs{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Songs{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
