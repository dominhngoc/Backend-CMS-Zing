package models

import (
	"fmt"
	. "github.com/astaxie/beego/orm"
	"strconv"
)

type Artists struct {
	Id        int    `json:",omitempty" json:",omitempty"`
	Fullname  string `valid:"Required;MinSize(6);MaxSize(50)" json:",omitempty"`
	Birthday  string `valid:"Required;" json:",omitempty"`
	AvatarUrl string `valid:"Required" json:",omitempty"`
	Followers int    `valid:"Required" json:",omitempty"`
	Country   string `json:",omitempty"`
	Biography string `json:",omitempty"`
	SongName  string `json:",omitempty"`
}
type ArtistsFormated struct {
	Id        int      `json:",omitempty"`
	Fullname  string   `json:",omitempty"`
	Birthday  string   `json:",omitempty"`
	AvatarUrl string   `json:",omitempty"`
	Followers int      `json:",omitempty"`
	Country   string   `json:",omitempty"`
	Biography string   `json:",omitempty"`
	Songs     []string `json:",omitempty"`
}
type ArtistsSwagger struct {
	Id        int
	Followers int
	Fullname  string
	Birthday  string
	AvatarUrl string
	Country   string
	Biography string
	Songs     []string
}

func (t *Artists) TableName() string {
	return "artists"
}

func init() {
	RegisterModel(new(Artists))
}
func ArtistCopy(artist1 *Artists, artist2 *ArtistsFormated) {
	artist2.Id = artist1.Id
	artist2.Fullname = artist1.Fullname
	artist2.Birthday = artist1.Birthday
	artist2.AvatarUrl = artist1.AvatarUrl
	artist2.Followers = artist1.Followers
	artist2.Country = artist1.Country
	artist2.Biography = artist1.Biography
	artist2.Songs = append(artist2.Songs, artist1.SongName)
}

func FormatArtistList(list1 []Artists) (list2 []ArtistsFormated) {
	newArtist := ArtistsFormated{}
	j := 0
	list2 = append(list2, newArtist)
	ArtistCopy(&list1[0], &list2[j])
	for i := 1; i < len(list1); i++ {
		if list1[i].Id == list1[i-1].Id {
			if list1[i].SongName != "" {
				list2[j].Songs = append(list2[j].Songs, list1[i].SongName)
			}
		} else {
			j++
			list2 = append(list2, newArtist)
			ArtistCopy(&list1[i], &list2[j])
		}
	}
	return list2
}

// AddArtists insert a new Artists into database and returns
// last inserted Id on success.
func AddArtists(m *Artists) (id int64, err error) {
	qb, _ := NewQueryBuilder("mysql")
	qb.InsertInto("artists",
		"fullname", "followers", "avatar_url", "birthday", "country", "biography").
		Values("?", "?", "?", "?", "?", "?")
	sql := qb.String()
	o := NewOrm()
	res, err := o.Raw(sql, m.Fullname, m.Followers, m.AvatarUrl, m.Birthday, m.Country, m.Biography).Exec()
	if err != nil {
		return -1, err
	}
	id, _ = res.LastInsertId()
	return id,nil
}

// GetArtistsById retrieves Artists by Id. Returns error if
// Id doesn't exist
//func GetArtistsById(id int) (v *Artists, err error) {
//	o := orm.NewOrm()
//	v = &Artists{Id: id}
//	if err = o.Read(v); err == nil {
//		return v, nil
//	}
//	return nil, err
//}
func GetArtistsById(id int) (list2 []ArtistsFormated, err error) {
	var list1 []Artists
	qb, _ := NewQueryBuilder("mysql")
	// Construct query object
	qb.Select("artists.id as id",
		"artists.fullname as fullname",
		"artists.birthday",
		"artists.avatar_url",
		"artists.followers",
		"artists.country",
		"artists.biography",
		"songs.name as song_name",
	)
	qb.From("artists").
		LeftJoin("artist_song").On("artists.id=artist_song.artist_id").
		LeftJoin("songs").On("artist_song.song_id=songs.id")
	sql := qb.String()
	sql += " WHERE artists.is_deleted = 0 AND artists.id = " + strconv.Itoa(id)
	// execute the raw query string
	fmt.Println(sql)
	o := NewOrm()
	o.Raw(sql).QueryRows(&list1)
	if list1 != nil {
		list2 = FormatArtistList(list1)
	}
	fmt.Println(list2)
	return list2, err
}

// GetAllArtists retrieves all Artists matches certain condition. Returns empty list if
func GetAllArtists(fields []string, searchByName string) (list2 []ArtistsFormated, err error) {
	var list1 []Artists
	qb, _ := NewQueryBuilder("mysql")
	// Construct query object
	if fields != nil {
		// Reformat field string to exec sql
		for i := 0; i < len(fields); i++ {
			fields[i] = "artists." + fields[i] + " as " + fields[i]
		}
		fields = append(fields, "artists.id as id")
		qb.Select(fields...)
	}
	if fields == nil { // query all
		qb.Select("artists.id as id",
			"artists.fullname as fullname",
			"artists.birthday",
			"artists.avatar_url",
			"artists.followers",
			"artists.country",
			"artists.biography",
			"songs.name as song_name",
		)
	}
	qb.From("artists").
		LeftJoin("artist_song").On("artists.id=artist_song.artist_id").
		LeftJoin("songs").On("artist_song.song_id=songs.id")
	sql := qb.String()
	sql += " WHERE artists.is_deleted = 0"
	// find name
	if searchByName != "" {
		sql += " AND fullname LIKE '%" + searchByName + "%' "
	}
	sql += " ORDER BY artists.id "
	fmt.Println(sql)
	// execute the raw query string
	o := NewOrm()
	o.Raw(sql).QueryRows(&list1)
	fmt.Println(list1)
	//Format artists list
	if list1 != nil {
		list2 = FormatArtistList(list1)
	}
	return list2, err
}

// UpdateArtists updates Artists by Id and returns error if
// the record to be updated doesn't exist
func UpdateArtistsById(m *Artists) (err error) {
	o := NewOrm()
	v := Artists{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteArtists deletes Artists by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArtists(id int, o Ormer) (err error) {
	sql := "UPDATE artists SET is_deleted = 1 WHERE artists.id = " + strconv.Itoa(id)
	_, err = o.Raw(sql).Exec()
	return err
}

func DeleteTransaction(id int) (err error) {
	o := NewOrm()
	err = o.Begin()
	// delete artist
	err1 := DeleteArtists(id, o)
	if err1 != nil {
		err = o.Rollback()
		fmt.Println(err1)
	}
	// delete artist_song
	err2 := DeleteArtistSong(id, o)
	if err2 != nil {
		err = o.Rollback()
		fmt.Println(err2)
	}
	err = o.Commit()
	return
}
