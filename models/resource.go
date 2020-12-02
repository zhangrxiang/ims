package models

import (
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
	model := db.DB.Where(&r).
		Order("rh_id DESC").
		Order("id DESC").
		Find(&resources)
	return resources, model.Error
}

func (r *ResourceModel) First() (*ResourceModel, error) {
	resource := &ResourceModel{}
	model := db.DB.First(resource, r.ID)
	return resource, model.Error
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

func (r *ResourceModel) All() ([]ResourceModel, error) {
	var resources []ResourceModel
	model := db.DB.Order("id DESC").Find(&resources)
	return resources, model.Error
}

func (r *ResourceModel) Insert() error {
	return db.DB.Create(r).Error
}

//更新
func (r *ResourceModel) Update() error {
	resource := &ResourceModel{}
	return db.DB.Model(resource).Updates(r).Error
}

func (r *ResourceModel) DeleteBy() error {
	return db.DB.Model(r).Delete(r).Error
}
