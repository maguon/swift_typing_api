package models

import "time"

type UserDevice struct {
	Id            int `gorm:"primaryKey;column:id"`
	AppType       int `binding:"required"`
	AppVersion    string
	AppVersionNum int    `binding:"required"`
	DeviceType    int    `binding:"required"`
	DeviceId      string `binding:"required"`
	DeviceToken   string
	Status        int `gorm:"default:1"`
	UserId        int
}

type UserDeviceOut struct {
	Id            int `gorm:"primaryKey;column:id"`
	AppType       int
	AppVersion    string
	AppVersionNum int
	DeviceType    int
	DeviceId      string
	DeviceToken   string
	Status        int
	UserId        int
	CreatedOn     time.Time
	UpdatedOn     time.Time
}

type UserDeviceQuery struct {
	Id            int    `json:"id" gorm:"primaryKey;column:id"`
	AppType       int    `json:"app_type" form:"app_type,default=-1"`
	DeviceType    int    `json:"device_type" form:"device_type,default=-1"`
	DeviceId      string `json:"device_id" form:"device_id"`
	DeviceToken   string `json:"device_token" form:"device_token"`
	UserId        int    `json:"user_id" form:"user_id,default=-1"`
	AppVersionNum int    `json:"app_version_num" form:"app_version_num,default=-1"`
	Status        int    `json:"status" form:"status,default=-1"`
	Start         int    `json:"start" form:"start,default=-1"`
	Size          int    `json:"size" form:"size,default=-1"`
}
