package model

import (
	"gorm.io/gorm"
	"time"
)

type Reserve struct {
	gorm.Model
	//学生id
	StudentId uint `json:"student_id" gorm:"student_id"`
	//座位id
	SeatId uint `json:"seat_id" gorm:"seat_id"`
	//预约开始时间
	AppointmentTimeStart time.Time `json:"appointment_time_start" gorm:"appointment_time_start"`
	//预约结束时间
	AppointmentTimeEnd time.Time `json:"appointment_time_end" gorm:"appointment_time_end"`
	//预约状态(0-未开始，1-使用中，2-已结束)
	Status int `json:"status" gorm:"status"`
}
