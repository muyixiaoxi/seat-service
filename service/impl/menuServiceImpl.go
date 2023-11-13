package service

import (
	"seat-service/model"
)

type MenuService struct {
}

func (*MenuService) GetMenu() (menus []*model.Menu) {
	db.Where("parent_id IS NULL").Preload("Children").Find(&menus)
	return
}
