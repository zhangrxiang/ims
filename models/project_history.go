package models

import "time"

type ProjectHistoryModel struct {
	ID          int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ProjectId   int       `json:"project_id" gorm:"not null"`
	Version     string    `json:"version" gorm:"not null"`
	ResourceIds string    `json:"resource_ids"`
	Log         string    `json:"log"`
	Hash        string    `json:"hash"`
	Download    int       `json:"download"`
	CreatedAt   time.Time `json:"created_at"`
}

func (ph *ProjectHistoryModel) Insert() (*ProjectHistoryModel, error) {
	model := db.DB.Create(ph)
	return model.Value.(*ProjectHistoryModel), model.Error
}
