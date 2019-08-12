package models

import (
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	ID       int    `json:"id",gorm:"primary_key;AUTO_INCREMENT"`
	Username string `json:"username",gorm:"not null;unique;type:varchar(30)"`
	Password string `json:"password",gorm:"not null;type:varchar(20)"`
	Role     string `json:"role"`
	Phone    string `json:"phone",gorm:"not null"`
	Mail     string `json:"mail",gorm:"not null"`
}

func (u *UserModel) Find() (*UserModel, error) {
	model := db.DB.Where("username = ? AND password = ?", u.Username, u.Password).Find(u)
	return model.Value.(*UserModel), model.Error
}

func (u *UserModel) FindByUsername() (*UserModel, error) {
	model := db.DB.Where("username = ?", u.Username).Find(u)
	return model.Value.(*UserModel), model.Error
}

func (u *UserModel) FindByID() (*UserModel, error) {
	model := db.DB.First(u, u.ID)
	return u, model.Error
}

func (u *UserModel) All() (*[]UserModel, error) {
	var users []UserModel
	model := db.DB.Find(&users)
	return model.Value.(*[]UserModel), model.Error
}

func (u *UserModel) Insert() (*UserModel, error) {
	_, err := u.FindByUsername()
	if err == gorm.ErrRecordNotFound {
		model := db.DB.Create(u)
		return model.Value.(*UserModel), model.Error
	}
	return nil, RecordExists
}

func (u *UserModel) Delete(ids []int) (*UserModel, error) {
	model := db.DB.Where(ids).Delete(u)
	if model.RowsAffected == 0 {
		return nil, NoRecordExists
	}
	return model.Value.(*UserModel), model.Error
}

func (u *UserModel) Update() (*UserModel, error) {
	model := db.DB.Model(u).Updates(u)
	if model.RowsAffected == 0 {
		return nil, NoRecordExists
	}
	return model.Value.(*UserModel), model.Error
}
