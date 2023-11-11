package model

import "gorm.io/gorm"

type Campus struct {
	gorm.Model
	//校区(0-西校区，1-东区，2-新东区)
	Name string `json:"name" gorm:"name"`
}
