package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sync"
)

var db *DB
var dbOnce sync.Once

type DB struct {
	DB *gorm.DB
}

//唯一冲突
const UniqueFailed = "UNIQUE constraint failed"

func GetDBInstance() *DB {
	dbOnce.Do(func() {
		db = &DB{}
		db.init()
	})
	return db
}

func (db *DB) init() {
	var err error
	db.DB, err = gorm.Open("sqlite3", "./models/model.db")
	if err != nil {
		panic("failed to connect database")
	}
	//db.DB.AutoMigrate(&models.UserModel{})
}

func (db *DB) Close() {
	err := db.DB.Close()
	if err != nil {
		panic("failed to close database")
	}
}

func (db *DB) GetDB() *gorm.DB {
	return db.DB
}
