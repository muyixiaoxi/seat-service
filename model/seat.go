package model

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	//校区id
	CampusId uint `json:"campus_id" gorm:"Campus_id"`
	//教学楼id
	BuildingId uint `json:"building_id" gorm:"building_id"`
	//教室id
	ClassroomId uint `json:"classroom_id" gorm:"classroom_id"`
	//预约状态(0-空闲，1-已预约，2-正在使用)
	Reserve int `json:"reserve" gorm:"reserve"`
	//状态(0-关闭，1-正常)
	Status int `json:"status" gorm:"status"`
}
