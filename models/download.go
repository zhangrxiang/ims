package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DownloadModel struct {
	ID         int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId     int       `json:"user_id"`
	ResourceId int       `json:"resource_id"`
	CreateAt   time.Time `json:"create_at"`
}

func (dm *DownloadModel) Find() (*DownloadModel, error) {
	downloadModel := &DownloadModel{}
	model := db.DB.First(&downloadModel, dm.ID)
	return downloadModel, model.Error
}

func (dm *DownloadModel) Insert() (*DownloadModel, error) {
	model := db.DB.Create(dm)
	return model.Value.(*DownloadModel), model.Error
}

func (dm *DownloadModel) Increment() (*DownloadModel, error) {
	model := db.DB.Model(dm).Update("price", gorm.Expr("price  + 1"))
	return model.Value.(*DownloadModel), model.Error
}
