package model

import "gorm.io/gorm"

type Dictionary struct {
	gorm.Model
	//名称
	Name string `json:"name" gorm:"name"`
	//代码
	Code string `json:"field_name" gorm:"field_name"`
	//状态(0-关闭，1-启用)
	DictionaryStatus string `json:"dictionary_status" gorm:"dictionary_status"`
}

type DictionarySlave struct {
	gorm.Model
	//名称
	Name string `json:"name" gorm:"name"`
	//值
	Value string `json:"value" gorm:"value"`
	//状态(0-关闭，1-启用)
	DictionarySlaveStatus string `json:"dictionary_slave_status" gorm:"dictionary_slave_status"`
	//master id
	DictionaryId uint `json:"dictionary_id" gorm:"dictionary_id"`
}
