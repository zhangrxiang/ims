package models

import (
	"time"
)

type DownloadModel struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId    int       `json:"user_id"`
	RHId      int       `json:"rh_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (dm *DownloadModel) FirstBy() (*DownloadModel, error) {
	downloadModel := &DownloadModel{}
	model := db.DB.Where(dm).First(&downloadModel)
	return downloadModel, model.Error
}

func (dm *DownloadModel) Insert() (*DownloadModel, error) {
	model := db.DB.Create(dm)
	return model.Value.(*DownloadModel), model.Error
}
