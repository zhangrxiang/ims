package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"simple-ims/utils"
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
		(*LogModel)(nil),
		(*ResourceModel)(nil),
		(*ResourceTypeModel)(nil),
		(*ResourceHistoryModel)(nil),
		(*DownloadModel)(nil),
		(*ProjectModel)(nil),
		(*ProjectHistoryModel)(nil),
	} {
		if !db.DB.HasTable(v) {
			db.DB.CreateTable(v)
			switch v.(type) {
			case *UserModel:
				for _, v := range []*UserModel{{
					ID:       1,
					Username: "admin",
					Password: utils.Encode("123456"),
					Role:     Admin,
				}, {
					ID:       2,
					Username: "atian",
					Password: utils.Encode("123456"),
					Role:     Downloader,
				}} {
					db.DB.Create(v)
				}
			case *ResourceTypeModel:
				for _, v := range []*ResourceTypeModel{{
					ID:   1,
					Name: "震动",
					Desc: "关于震动",
				}, {
					ID:   2,
					Name: "周界",
					Desc: "关于周界",
				}, {
					ID:   3,
					Name: "测温",
					Desc: "关于测温",
				}} {
					db.DB.Create(v)
				}
			}
		}
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
