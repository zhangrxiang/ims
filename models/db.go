package models

import (
	"errors"
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
var (
	RecordExists   = errors.New("数据记录已经存在")
	NoRecordExists = errors.New("无数据记录存在")
)

func GetDBInstance() *DB {
	dbOnce.Do(func() {
		db = &DB{}
		db.Init()
	})
	return db
}

func (db *DB) Init() {
	var err error
	db.DB, err = gorm.Open("sqlite3", "./database.db")
	if err != nil {
		panic("failed to connect database")
	}

	for _, v := range []interface{}{
		(*UserModel)(nil),
		(*ResourceModel)(nil),
		(*ResourceTypeModel)(nil),
		(*ResourceHistoryModel)(nil),
	} {
		db.DB.AutoMigrate(v)
	}

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
