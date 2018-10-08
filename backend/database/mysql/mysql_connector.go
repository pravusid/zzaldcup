package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang-server/model"
	"os"
)

var db *gorm.DB

func init() {
	var err error

	if db, err = gorm.Open("mysql", os.Getenv("DATABASE")); err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(true)

	if os.Getenv("PROFILE") == "dev" {
		autoCreateTables()
		autoMigrateTables()
	}
}

func DB() *gorm.DB {
	return db
}

func autoCreateTables() {
	if !db.HasTable(&model.Match{}) {
		db.CreateTable(&model.Match{})
	}
	if !db.HasTable(&model.Competitor{}) {
		db.CreateTable(&model.Competitor{})
	}
}

func autoMigrateTables() {
	db.AutoMigrate(&model.Match{})
	db.AutoMigrate(&model.Competitor{})
}

func autoDropTables() {
	db.DropTableIfExists(&model.Match{}, &model.Match{})
	db.DropTableIfExists(&model.Competitor{}, &model.Competitor{})
}
