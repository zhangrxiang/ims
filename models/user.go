package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

const (
	Admin      = "admin"
	Uploader   = "uploader"
	Downloader = "downloader"
)

type UserModel struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username  string    `json:"username" gorm:"not null;unique;type:varchar(30)"`
	Password  string    `json:"password" gorm:"not null;type:varchar(20)"`
	Role      string    `json:"role"`
	Phone     string    `json:"phone" gorm:"not null"`
	Mail      string    `json:"mail" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var ErrUserNameExists = errors.New("用户名不能重名")

func (u *UserModel) BeforeCreate(*gorm.DB) (err error) {
	u2 := &UserModel{Username: u.Username}
	if _, err := u2.First(); err == nil {
		return ErrUserNameExists
	}
	return
}

func (u *UserModel) Insert() error {
	return db.DB.Create(u).Error
}

func (u *UserModel) First() (*UserModel, error) {
	var users UserModel
	m := db.DB.Model(u).Where(u).First(&users)
	return &users, m.Error
}

func (u *UserModel) Find() ([]UserModel, error) {
	var users []UserModel
	model := db.DB.Model(u).Where(u).Order("id DESC").Find(&users)
	return users, model.Error
}

func (u *UserModel) Delete() error {
	return db.DB.Model(u).Delete(u).Error
}

func (u *UserModel) BeforeUpdate(*gorm.DB) (err error) {
	u2 := &UserModel{Username: u.Username}
	if u2, err = u2.First(); err != nil {
		return nil
	}
	if u2.ID != u.ID {
		return ErrUserNameExists
	}
	return
}

func (u *UserModel) Update() error {
	return db.DB.Model(u).
		Updates(u).
		UpdateColumn("phone", u.Phone).
		UpdateColumn("mail", u.Mail).
		Error
}
