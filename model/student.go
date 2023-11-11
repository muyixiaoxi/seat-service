package model

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	//头像
	Head string `json:"head" gorm:"head"`
	//学院
	College string `json:"college" gorm:"college"`
	//年级
	Grade string `json:"grade" gorm:"grade"`
	//班级
	Class string `json:"class" gorm:"class"`
	//姓名
	Name string `json:"name" gorm:"name"`
	//性别(0-男，1-女，3-未知)
	Gender int `json:"gender" gorm:"gender"`
	//学号
	StudentNumber int `json:"student_number" gorm:"student_number"`
	//状态(0-正常，1-违规)
	Status int `json:"status" gorm:"status"`
	//违规锁定开始时间
	ViolationLockTimeStart time.Time `json:"violation_lock_time_start" gorm:"violation_lock_time_start"`
	//违规锁定结束时间
	ViolationLockTimeEnd time.Time `json:"violation_lock_time_end" gorm:"violation_lock_time_end"`
}
