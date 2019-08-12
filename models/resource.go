package models

import (
	"errors"
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
	ID       int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name     string    `json:"name" gorm:"not null"`
	Type     int       `json:"type" gorm:"not null"`
	File     string    `json:"file"`
	Path     string    `json:"path"`
	Hash     string    `json:"hash"`
	Version  string    `json:"version" gorm:"not null"`
	Desc     string    `json:"desc" gorm:"not null"`
	CreateAt time.Time `json:"create_at"`
}

func (r *ResourceModel) FindByHash(h string) (*ResourceModel, error) {
	db := utils.GetDBInstance().DB
	db.AutoMigrate(r)
	model := db.Where("hash = ?", h).First(r)
	if model.RowsAffected == 0 {
		return nil, nil
	}
	return model.Value.(*ResourceModel), model.Error
}

func (r *ResourceModel) All() (*[]ResourceModel, error) {
	db := utils.GetDBInstance().DB
	db.AutoMigrate(r)
	var resources []ResourceModel
	model := db.Order("id DESC").Find(&resources)

	return model.Value.(*[]ResourceModel), model.Error
}

func (r *ResourceModel) Insert() (*ResourceModel, error) {
	db := utils.GetDBInstance().DB
	db.AutoMigrate(r)
	model := db.Create(r)

	return model.Value.(*ResourceModel), model.Error
}

//根据ID删除
func (r *ResourceModel) DeleteByIds(ids []int) (*ResourceModel, error) {
	db := utils.GetDBInstance().DB
	model := db.Where(ids).Delete(r)

	if model.RowsAffected == 0 {
		return nil, errors.New("无此资源")
	}

	return model.Value.(*ResourceModel), model.Error
}

//更新
func (r *ResourceModel) Update() (*ResourceModel, error) {
	db := utils.GetDBInstance().DB
	model := db.Model(r).Updates(r)

	if model.RowsAffected == 0 {
		return nil, errors.New("无此资源")
	}
	return model.Value.(*ResourceModel), model.Error
}
