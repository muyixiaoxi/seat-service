package service

import (
	"seat-service/initialization"
	"seat-service/model"
)

var db = initialization.DB

type UserService struct {
}

func (u *UserService) Test(i int) int {
	return i + 1
}

func (u *UserService) Login(user *model.User) (err error) {
	err = db.First(u).Error
	return err
}
