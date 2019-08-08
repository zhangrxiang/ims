package models

import "simple-ims/utils"

type UserModel struct {
	ID       int    `json:"id",gorm:"primary_key;AUTO_INCREMENT"`
	Username string `json:"username",gorm:"not null;unique;type:varchar(30)"`
	Password string `json:"password",gorm:"not null;type:varchar(20)"`
	Role     string `json:"role"`
	Phone    string `json:"phone",gorm:"not null"`
	Mail     string `json:"mail",gorm:"not null"`
}

func (u *UserModel) Find() (*UserModel, error) {

	db := utils.GetDBInstance().DB
	model := db.Where("username = ? AND password >= ?", u.Username, u.Password).Find(u)
	return model.Value.(*UserModel), model.Error

}

func (u *UserModel) All() (*[]UserModel, error) {
	db := utils.GetDBInstance().DB
	var users []UserModel
	model := db.Find(&users)
	return model.Value.(*[]UserModel), model.Error
}

func (u *UserModel) FindByID() (*UserModel, error) {
	db := utils.GetDBInstance().DB

	db = db.First(u, u.ID)

	return u, db.Error
}

func (u *UserModel) Insert() (*UserModel, error) {

	db := utils.GetDBInstance().DB
	db.AutoMigrate(u)

	db = db.Create(u)

	return db.Value.(*UserModel), db.Error
}
