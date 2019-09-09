package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type SongsTable_20190905_094008 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &SongsTable_20190905_094008{}
	m.Created = "20190905_094008"

	migration.Register("SongsTable_20190905_094008", m)
}

// Run the migrations
func (m *SongsTable_20190905_094008) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("DROP TABLE IF EXISTS `songs`")

	m.CreateTable("songs", "InnoDB", "utf8mb4")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.NewCol("name").SetDataType("VARCHAR(255)").SetNullable(false)
	m.NewCol("lyrics").SetDataType("TEXT").SetNullable(true)
	m.NewCol("music_url").SetDataType("VARCHAR(255)").SetNullable(false)
	m.NewCol("kind").SetDataType("VARCHAR(255)").SetNullable(false)
	m.NewCol("image_url").SetDataType("VARCHAR(255)").SetNullable(true)
	m.NewCol("lrc_url").SetDataType("VARCHAR(255)").SetNullable(true)
	m.NewCol("released").SetDataType("VARCHAR(255)").SetNullable(true)
	m.NewCol("created_user_id").SetDataType("INT(10)").SetNullable(false).SetDefault("1")
	m.NewCol("updated_user_id").SetDataType("INT(10)").SetNullable(false).SetDefault("1")
	m.NewCol("is_deleted").SetDataType("INT(1)").SetNullable(false).SetDefault("0")
	m.NewCol("created_at").SetDataType("TIMESTAMP").SetNullable(true).SetDefault("CURRENT_TIMESTAMP")
	m.NewCol("updated_at").SetDataType("TIMESTAMP").SetNullable(true).SetDefault("CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP")

	sql := m.GetSQL()
	m.SQL(sql)
}

// Reverse the migrations
func (m *SongsTable_20190905_094008) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `songs`")


}
