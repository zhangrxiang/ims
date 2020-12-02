package models

import (
	"gorm.io/gorm"
	"time"
)

type ResourceHistoryModel struct {
	ID         int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ResourceID int       `json:"resource_id" gorm:"not null"`
	UserId     int       `json:"user_id"`
	Version    string    `json:"version" gorm:"not null"`
	Log        string    `json:"log"`
	File       string    `json:"file"`
	Path       string    `json:"path"`
	Hash       string    `json:"hash"`
	Download   int       `json:"download"`
	CreatedAt  time.Time `json:"created_at"`
}

func (rh *ResourceHistoryModel) Increment() error {
	return db.DB.Model(rh).Update("download", gorm.Expr("download  + 1")).Error
}

func (rh *ResourceHistoryModel) Insert() error {
	return db.DB.Create(rh).Error
}

//单数据查询
func (rh *ResourceHistoryModel) FirstBy() (*ResourceHistoryModel, error) {
	var resource ResourceHistoryModel
	model := db.DB.Where(rh).Order("id DESC").First(&resource)
	if model.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &resource, model.Error
}

//多数据查询
func (rh *ResourceHistoryModel) FindBy() ([]ResourceHistoryModel, error) {
	var resources []ResourceHistoryModel
	model := db.DB.Order("id DESC").Find(&resources, "resource_id = ?", rh.ResourceID)
	return resources, model.Error
}

func (rh *ResourceHistoryModel) FindByIDs(ids []int) ([]ResourceHistoryModel, error) {
	var resources []ResourceHistoryModel
	model := db.DB.Order("id DESC").Where(ids).Find(&resources)
	return resources, model.Error
}

func (rh *ResourceHistoryModel) FindValueBy(key string) ([]interface{}, error) {
	var values []interface{}
	model := db.DB.Model(rh).Pluck(key, &values)
	return values, model.Error
}

func (rh *ResourceHistoryModel) FindIDBy() ([]int, error) {
	var values []int
	model := db.DB.Model(rh).Where(rh).Pluck("id", &values)
	return values, model.Error
}

//更新
func (rh *ResourceHistoryModel) Update() error {
	resource := &ResourceHistoryModel{}
	return db.DB.Model(resource).Updates(rh).Error
}

func (rh *ResourceHistoryModel) DeleteBy(ids []int) error {
	model := db.DB.Model(rh).Where(ids).Delete(rh)
	return model.Error
}

func (rh *ResourceHistoryModel) Delete() error {
	model := db.DB.Model(rh).Delete(rh)
	return model.Error
}
