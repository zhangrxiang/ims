package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ResourceModel struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId    int       `json:"user_id"`
	RHId      int       `json:"rh_id"`
	Name      string    `json:"name" gorm:"not null"`
	Type      int       `json:"type" gorm:"not null"`
	Desc      string    `json:"desc" gorm:"not null"`
	Download  int       `json:"download"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *ResourceModel) FindBy() ([]ResourceModel, error) {
	var resources []ResourceModel
	model := db.DB.Where(&r).Order("id DESC").Find(&resources)
	return resources, model.Error
}

func (r *ResourceModel) First() (*ResourceModel, error) {
	resource := &ResourceModel{}
	model := db.DB.First(resource, r.ID)
	return resource, model.Error
}

func (r *ResourceModel) FindByIds(ids []int) (*[]ResourceModel, error) {
	var resources []ResourceModel
	model := db.DB.Where(ids).Order("id DESC").Find(&resources)
	if model.RowsAffected == 0 {
		return nil, NoRecordExists
	}
	return model.Value.(*[]ResourceModel), model.Error
}

func (r *ResourceModel) FirstByHash(h string) (*ResourceModel, error) {
	resource := &ResourceModel{}
	model := db.DB.Where("hash = ?", h).First(resource)
	if model.RowsAffected == 0 {
		return nil, NoRecordExists
	}
	return resource, model.Error
}

func (r *ResourceModel) FindByType(t int) ([]ResourceModel, error) {
	var resources []ResourceModel
	model := db.DB.Where("type = ?", t).Find(&resources)
	return resources, model.Error
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

func (r *ResourceModel) Increment() (*ResourceModel, error) {
	model := db.DB.Model(r).Update("download", gorm.Expr("download + 1"))
	return model.Value.(*ResourceModel), model.Error
}

func (r *ResourceModel) DeleteBy() error {
	model := db.DB.Model(r).Delete(r)
	return model.Error
}
