package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type ResourceModel struct {
	Id        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId    int       `json:"user_id"`
	RHId      int       `json:"rh_id"`
	Name      string    `json:"name" gorm:"not null"`
	Type      int       `json:"type" gorm:"not null"`
	Desc      string    `json:"desc" gorm:"not null"`
	Download  int       `json:"download"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var ErrRMNameExists = errors.New("资源不能重名")

func (r *ResourceModel) BeforeCreate(*gorm.DB) (err error) {
	r2 := &ResourceModel{Name: r.Name}
	if _, err := r2.First(); err == nil {
		return ErrRMNameExists
	}
	return
}

func (r *ResourceModel) Insert() error {
	return db.DB.Create(r).Error
}

func (r *ResourceModel) First() (*ResourceModel, error) {
	resource := &ResourceModel{}
	model := db.DB.Where(r).First(resource)
	return resource, model.Error
}

func (r *ResourceModel) Find() ([]ResourceModel, error) {
	var resources []ResourceModel
	m := db.DB.Model(r)
	if r.Type != 0 {
		m.Where("type = ?", r.Type)
	}
	if r.Name != "" {
		m.Where("name = ?", r.Name)
	}
	m.Order("updated_at DESC").
		Order("id DESC").
		Find(&resources)
	return resources, m.Error
}

func (r *ResourceModel) BeforeUpdate(*gorm.DB) (err error) {
	if _, err = (&ResourceModel{Id: r.Id}).First(); err != nil {
		return err
	}
	if r.Name == "" {
		return
	}
	r2 := &ResourceModel{Name: r.Name}
	if r2, err = r2.First(); err != nil {
		return nil
	}
	if r.Id != r2.Id {
		return ErrRMNameExists
	}
	return
}

//更新
func (r *ResourceModel) Update() error {
	return db.DB.Model(r).Updates(r).Error
}

func (r *ResourceModel) Delete() error {
	return db.DB.Model(r).Delete(r).Error
}
