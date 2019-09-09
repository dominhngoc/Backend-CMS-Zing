package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ArtistSongTable_20190828_110326 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ArtistSongTable_20190828_110326{}
	m.Created = "20190828_110326"

	migration.Register("ArtistSongTable_20190828_110326", m)
}

// Run the migrations
func (m *ArtistSongTable_20190828_110326) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("DROP TABLE IF EXISTS `artist_song`")

	m.CreateTable("artist_song", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.NewCol("artist_id").SetDataType("INT(10)").SetNullable(false)
	m.NewCol("song_id").SetDataType("INT(10)").SetNullable(false)
	m.NewCol("created_at").SetDataType("TIMESTAMP").SetNullable(true).SetDefault("CURRENT_TIMESTAMP")
	m.NewCol("updated_at").SetDataType("TIMESTAMP").SetNullable(true).SetDefault("CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP")
	sql := m.GetSQL()
	m.SQL(sql)
}

// Reverse the migrations
func (m *ArtistSongTable_20190828_110326) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `artist_song`")
}
