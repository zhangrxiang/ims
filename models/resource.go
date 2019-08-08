package models

import (
	"simple-ims/utils"
	"time"
)

type ResourceType byte

const (
	COMMON ResourceType = iota //通用
	RSMS                       //周界
	DFVS                       //测温
	RFVS                       //振动
)

func (rt ResourceType) String() string {
	switch rt {
	case COMMON:
		return "通用"
	case RSMS:
		return "周界"
	case DFVS:
		return "测温"
	case RFVS:
		return "振动"
	default:
		return "未知"
	}
}

type ResourceModel struct {
	ID       int          `json:"id",gorm:"primary_key;AUTO_INCREMENT"`
	Name     string       `json:"name",gorm:"not null"`
	Type     ResourceType `json:"type",gorm:"not null"`
	File     string       `json:"file"`
	Version  string       `json:"version",gorm:"not null"`
	Desc     string       `json:"desc",gorm:"not null"`
	CreateAt time.Time    `json:"create_at"`
}

func (r *ResourceModel) All() (*[]ResourceModel, error) {
	db := utils.GetDBInstance().DB
	var resources []ResourceModel
	model := db.Find(&resources)
	return model.Value.(*[]ResourceModel), model.Error
}

func (r *ResourceModel) Insert() (*UserModel, error) {

	db := utils.GetDBInstance().DB
	db.AutoMigrate(r)
	db = db.Create(r)

	return db.Value.(*UserModel), db.Error
}
