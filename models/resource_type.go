package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type ResourceTypeModel struct {
	Id        int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	Desc      string    `json:"desc" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var ErrRTMNameExists = errors.New("资源类型不能重名")

func (r *ResourceTypeModel) BeforeCreate(*gorm.DB) (err error) {
	r2 := &ResourceTypeModel{Name: r.Name}
	if _, err := r2.First(); err == nil {
		return ErrRTMNameExists
	}
	return
}

func (r *ResourceTypeModel) Insert() error {
	return db.DB.Create(r).Error
}

func (r *ResourceTypeModel) First() (*ResourceTypeModel, error) {
	var rm ResourceTypeModel
	m := db.DB.Model(r).Where(r).First(&rm)
	return &rm, m.Error
}

//查询所有
func (r *ResourceTypeModel) Find() ([]ResourceTypeModel, error) {
	var resources []ResourceTypeModel
	model := db.DB.Order("id DESC").Find(&resources)
	return resources, model.Error
}

//根据ID删除
func (r *ResourceTypeModel) DeleteByIds(ids []int) error {
	return db.DB.Where(ids).Delete(r).Error
}

func (r *ResourceTypeModel) BeforeUpdate(*gorm.DB) (err error) {
	if _, err = r.First(); err != nil {
		return err
	}
	r2 := &ResourceTypeModel{Name: r.Name}
	if r2, err = r2.First(); err != nil {
		return nil
	}
	if r2.Id != r2.Id {
		return ErrRTMNameExists
	}
	return
}

func (r *ResourceTypeModel) Update() error {
	return db.DB.Model(r).Updates(r).Error
}
