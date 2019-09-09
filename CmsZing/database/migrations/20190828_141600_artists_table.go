package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ArtistsTable_20190828_141600 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ArtistsTable_20190828_141600{}
	m.Created = "20190828_141600"

	migration.Register("ArtistsTable_20190828_141600", m)
}

// Run the migrations
func (m *ArtistsTable_20190828_141600) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("DROP TABLE IF EXISTS `artists`")

	m.CreateTable("artists", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.NewCol("fullname").SetDataType("VARCHAR(255)").SetNullable(false)
	m.NewCol("birthday").SetDataType("VARCHAR(255)").SetNullable(false)
	m.NewCol("avatar_url").SetDataType("VARCHAR(255)").SetNullable(false)
	m.NewCol("followers").SetDataType("INT(10)").SetNullable(false)
	m.NewCol("country").SetDataType("VARCHAR(255)").SetNullable(true)
	m.NewCol("biography").SetDataType("TEXT").SetNullable(true)
	m.NewCol("created_user_id").SetDataType("INT(10)").SetNullable(false).SetDefault("1")
	m.NewCol("updated_user_id").SetDataType("INT(10)").SetNullable(false).SetDefault("1")
	m.NewCol("is_deleted").SetDataType("INT(1)").SetNullable(false).SetDefault("0")
	m.NewCol("created_at").SetDataType("TIMESTAMP").SetNullable(true).SetDefault("CURRENT_TIMESTAMP")
	m.NewCol("updated_at").SetDataType("TIMESTAMP").SetNullable(true).SetDefault("CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP")

	sql := m.GetSQL()
	m.SQL(sql)

}

// Reverse the migrations
func (m *ArtistsTable_20190828_141600) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `artists`")

}
