package models

import (
	"time"
)

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

func (r *ResourceModel) Find() (*ResourceModel, error) {
	resource := &ResourceModel{}
	model := db.DB.First(resource, r.ID)
	return resource, model.Error
}

func (r *ResourceModel) FindIds(ids []int) (*[]ResourceModel, error) {
	var resources []ResourceModel
	model := db.DB.Where(ids).Order("id DESC").Find(&resources)
	if model.RowsAffected == 0 {
		return nil, NoRecordExists
	}
	return model.Value.(*[]ResourceModel), model.Error
}

func (r *ResourceModel) FindByHash(h string) (*ResourceModel, error) {
	resource := &ResourceModel{}
	model := db.DB.Where("hash = ?", h).First(resource)
	if model.RowsAffected == 0 {
		return nil, NoRecordExists
	}
	return resource, model.Error
}

func (r *ResourceModel) All() (*[]ResourceModel, error) {
	var resources []ResourceModel
	model := db.DB.Order("id DESC").Find(&resources)
	return model.Value.(*[]ResourceModel), model.Error
}

func (r *ResourceModel) Insert() (*ResourceModel, error) {
	model := db.DB.Create(r)
	return model.Value.(*ResourceModel), model.Error
}

//根据ID删除
func (r *ResourceModel) DeleteByIds(ids []int) (*ResourceModel, error) {
	model := db.DB.Where(ids).Delete(r)
	if model.RowsAffected == 0 {
		return nil, NoRecordExists
	}
	return model.Value.(*ResourceModel), model.Error
}

//更新
func (r *ResourceModel) Update() (*ResourceModel, error) {
	resource := &ResourceModel{}
	model := db.DB.Model(resource).Updates(r)
	if model.RowsAffected == 0 {
		return nil, NoRecordExists
	}
	return model.Value.(*ResourceModel), model.Error
}
