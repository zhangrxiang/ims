package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type ProjectModel struct {
	Id        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	PHId      int       `json:"ph_id"`
	Name      string    `json:"name" gorm:"not null"`
	Desc      string    `json:"desc" gorm:"not null"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var ErrRroMNameExists = errors.New("项目不能重名")

func (p *ProjectModel) BeforeCreate(*gorm.DB) (err error) {
	p2 := &ProjectModel{Name: p.Name}
	if _, err := p2.First(); err == nil {
		return ErrRroMNameExists
	}
	return
}

func (p *ProjectModel) Insert() error {
	return db.DB.Create(p).Error
}

func (p *ProjectModel) Delete() error {
	return db.DB.Delete(p).Error
}

func (p *ProjectModel) First() (*ProjectModel, error) {
	var project ProjectModel
	model := db.DB.Model(p).Where(p).First(&project)
	return &project, model.Error
}

//多数据查询
func (p *ProjectModel) Find() ([]ProjectModel, error) {
	var projects []ProjectModel
	model := db.DB.Order("id DESC").Where(&p).Find(&projects)
	return projects, model.Error
}

func (p *ProjectModel) BeforeUpdate(*gorm.DB) (err error) {
	if _, err = (&ProjectModel{Id: p.Id}).First(); err != nil {
		return err
	}
	if p.Name == "" {
		return
	}
	p2 := &ProjectModel{Name: p.Name}
	if p2, err = p2.First(); err != nil {
		return nil
	}
	if p.Id != p2.Id {
		return ErrRroMNameExists
	}
	return
}

func (p *ProjectModel) Update() error {
	return db.DB.Model(p).Updates(p).Error
}
