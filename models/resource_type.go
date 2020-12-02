package models

import (
	"gorm.io/gorm"
	"time"
)

type ResourceTypeModel struct {
	ID        int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	Desc      string    `json:"desc" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *ResourceTypeModel) First() (*ResourceTypeModel, error) {
	t := &ResourceTypeModel{}
	model := db.DB.First(t, r.ID)
	return t, model.Error
}

//查询所有
func (r *ResourceTypeModel) All() ([]ResourceTypeModel, error) {
	var resources []ResourceTypeModel
	model := db.DB.Order("id DESC").Find(&resources)
	return resources, model.Error
}

func (r *ResourceTypeModel) FindByResourceName() (*ResourceTypeModel, error) {
	typeModel := &ResourceTypeModel{}
	model := db.DB.Where("name = ?", r.Name).Find(typeModel)
	return typeModel, model.Error
}

//根据ID删除
func (r *ResourceTypeModel) DeleteByIds(ids []int) error {
	return db.DB.Where(ids).Delete(r).Error

}

//更新
func (r *ResourceTypeModel) Update() error {
	typeModel, err := r.FindByResourceName()
	if err != gorm.ErrRecordNotFound && typeModel != nil {
		if typeModel.ID != r.ID {
			return RecordExists
		}
	}
	model := db.DB.Model(r).Updates(r)
	if model.RowsAffected == 0 {
		return NoRecordExists
	}
	return model.Error
}

//添加
func (r *ResourceTypeModel) Insert() error {
	_, err := r.FindByResourceName()
	if err == gorm.ErrRecordNotFound {
		return db.DB.Create(r).Error
	}
	return RecordExists
}
