package models

import (
	"gorm.io/gorm"
	"time"
)

type ProjectModel struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	PHId      int       `json:"ph_id"`
	Name      string    `json:"name" gorm:"not null"`
	Desc      string    `json:"desc" gorm:"not null"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *ProjectModel) Delete() error {
	return db.DB.Delete(p).Error
}

func (p *ProjectModel) FirstBy() (*ProjectModel, error) {
	var project ProjectModel
	model := db.DB.Order("id DESC").Where(&p).First(&project)
	if model.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &project, model.Error
}

//多数据查询
func (p *ProjectModel) FindBy() ([]ProjectModel, error) {
	var projects []ProjectModel
	model := db.DB.Order("id DESC").Where(&p).Find(&projects)
	return projects, model.Error
}

func (p *ProjectModel) Insert() error {
	return db.DB.Create(p).Error
}

func (p *ProjectModel) Update() error {
	return db.DB.Model(p).Updates(p).Error
}
