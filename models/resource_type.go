package models

import (
	"errors"
	"strings"
	"time"
)

type ResourceTypeModel struct {
	ID       int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name     string    `gorm:"unique;not null" json:"name"`
	Desc     string    `json:"desc" gorm:"not null"`
	CreateAt time.Time `json:"create_at"`
}

//查询所有
func (r *ResourceTypeModel) All() (*[]ResourceTypeModel, error) {

	db := GetDBInstance().DB
	var resources []ResourceTypeModel
	model := db.Order("id DESC").Find(&resources)
	return model.Value.(*[]ResourceTypeModel), model.Error

}

//根据ID删除
func (r *ResourceTypeModel) DeleteByIds(ids []int) (*ResourceTypeModel, error) {

	db := GetDBInstance().DB
	model := db.Where(ids).Delete(r)
	if model.RowsAffected == 0 {
		return nil, errors.New("无此资源分类")
	}
	return model.Value.(*ResourceTypeModel), model.Error

}

//更新
func (r *ResourceTypeModel) Update() (*ResourceTypeModel, error) {

	db := GetDBInstance().DB
	model := db.Model(r).Updates(r)

	if model.Error != nil && strings.Contains(model.Error.Error(), UniqueFailed) {
		return model.Value.(*ResourceTypeModel), errors.New("资源分类已经存在")
	}

	if model.RowsAffected == 0 {
		return nil, errors.New("无此资源分类")
	}
	return model.Value.(*ResourceTypeModel), model.Error

}

//添加
func (r *ResourceTypeModel) Insert() (*ResourceTypeModel, error) {

	db := GetDBInstance().DB
	model := db.Create(r)

	if model.Error != nil && strings.Contains(model.Error.Error(), UniqueFailed) {
		return model.Value.(*ResourceTypeModel), errors.New("资源分类已经存在")
	}

	return model.Value.(*ResourceTypeModel), model.Error
}
