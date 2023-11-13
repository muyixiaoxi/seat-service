package service

import "seat-service/model"

type Menu interface {
	GetMenu() model.Menu
}
