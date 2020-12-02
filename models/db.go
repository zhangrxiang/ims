package models

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	db.DB, err = gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrator := db.DB.Migrator()
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
		if !migrator.HasTable(v) {
			_ = migrator.CreateTable(v)
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
					Id:   1,
					Name: "震动",
					Desc: "关于震动",
				}, {
					Id:   2,
					Name: "周界",
					Desc: "关于周界",
				}, {
					Id:   3,
					Name: "测温",
					Desc: "关于测温",
				}} {
					db.DB.Create(v)
				}
			}
		}
		_ = db.DB.AutoMigrate(v)
	}
}

func (db *DB) Close() {
}

func (db *DB) GetDB() *gorm.DB {
	return db.DB
}
