package services

import "simple-ims/models"

type UserService struct{}

func (u *UserService) Add(model models.UserModel) (*models.UserModel, error) {
	return model.Insert()
}

func (u *UserService) Find(id int) (*models.UserModel, error) {
	return (&models.UserModel{ID: id}).Find()
}
