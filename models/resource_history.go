package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type ResourceHistoryModel struct {
	ID         int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ResourceID int       `json:"resource_id" gorm:"not null"`
	Version    string    `json:"version" gorm:"not null"`
	Log        string    `json:"log"`
	File       string    `json:"file"`
	Path       string    `json:"path"`
	Hash       string    `json:"hash"`
	Download   int       `json:"download"`
	CreateAt   time.Time `json:"create_at"`
}

func (rh *ResourceHistoryModel) Increment() (*ResourceHistoryModel, error) {
	log.Println(rh)
	model := db.DB.Model(rh).Update("download", gorm.Expr("download  + 1"))
	return nil, model.Error
}

func (rh *ResourceHistoryModel) Insert() (*ResourceHistoryModel, error) {
	model := db.DB.Create(rh)
	return model.Value.(*ResourceHistoryModel), model.Error
}

//单数据查询
func (rh *ResourceHistoryModel) FirstBy() (*ResourceHistoryModel, error) {
	var resource ResourceHistoryModel
	model := db.DB.Where(rh).First(&resource)
	return &resource, model.Error
}

//多数据查询
func (rh *ResourceHistoryModel) FindBy() (*[]ResourceHistoryModel, error) {
	var resources []ResourceHistoryModel
	model := db.DB.Order("id DESC").Find(&resources, "resource_id = ?", rh.ResourceID)
	return &resources, model.Error
}
