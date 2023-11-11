package model

import "gorm.io/gorm"

type Building struct {
	gorm.Model
	//教学楼名称(1号楼/弘德楼)
	Name string `json:"name" gorm:"name"`
	//所属校区id
	CampusId uint `json:"campus_id" gorm:"campus_id"`
	//状态(0-关闭，1-正常)
	Status int `json:"status" gorm:"status"`
}
