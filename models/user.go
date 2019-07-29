package models

import "simple-ims/utils"

type UserModel struct {
	ID       int    `json:"id",gorm:"primary_key;AUTO_INCREMENT"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
}

func (u *UserModel) Find() (*UserModel, error) {
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
