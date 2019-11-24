package database

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbConfig string = "root:ironman.abc@/mirrors?charset=utf8&parseTime=True&loc=Local"

var mirrorDB *gorm.DB = nil

func init() {

	db, err := gorm.Open("mysql", dbConfig)
	if err != nil {
		fmt.Println(err)
	}

	db.DB().Ping()

	mirrorDB = db
}

func GetDB() (db *gorm.DB, err error) {
	if mirrorDB == nil {
		return nil, errors.New("database is not open")
	}

	return mirrorDB, nil
}

func Close() {
	mirrorDB.Close()
}
