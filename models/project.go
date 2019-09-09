package models

import "time"

type ProjectModel struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	PHId      int       `json:"ph_id"`
	Name      string    `json:"name" gorm:"not null"`
	Desc      string    `json:"desc" gorm:"not null"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *ProjectModel) Insert() (*ProjectModel, error) {
	model := db.DB.Create(p)
	return model.Value.(*ProjectModel), model.Error
}
