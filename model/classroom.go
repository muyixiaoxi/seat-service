package model

import "gorm.io/gorm"

type Classroom struct {
	gorm.Model
	//校区id
	CampusId uint `json:"campus_id" gorm:"Campus_id"`
	//教学楼id
	BuildingId uint `json:"building_id" gorm:"building_id"`
	//教室位置
	Location string `json:"location" gorm:"location"`
	//座位总数
	Seating int `json:"seating" gorm:"seating"`
	//是否是合教(0-不是，1-是)
	IsLargeClassroom int `json:"is_large_classroom" gorm:"is_large_classroom"`
	//状态(0-关闭，1-正常)
	Status int `json:"status" gorm:"status"`
}
