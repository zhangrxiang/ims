package models

import "time"

type ProjectHistoryModel struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ProjectID int       `json:"project_id" gorm:"not null"`
	Version   string    `json:"version" gorm:"not null"`
	Log       string    `json:"log"`
	Hash      string    `json:"hash"`
	Download  int       `json:"download"`
	CreateAt  time.Time `json:"create_at"`
}

func (ph *ProjectHistoryModel) Insert() (*ProjectHistoryModel, error) {
	model := db.DB.Create(ph)
	return model.Value.(*ProjectHistoryModel), model.Error
}
