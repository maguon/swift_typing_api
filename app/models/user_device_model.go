package models

import "time"

type UserDevice struct {
	Id            int    `json:"id" gorm:"primaryKey;column:id"`
	AppType       int    `json:"appType" binding:"required"`
	AppVersion    string `json:"appVersion"`
	AppVersionNum int    `json:"appVersionNum" binding:"required"`
	DeviceType    int    `json:"deviceType" binding:"required"`
	DeviceId      string `json:"deviceId" binding:"required"`
	DeviceToken   string `json:"deviceToken"`
	Status        int    `gorm:"default:1"`
	UserId        int    `json:"userId"`
}

type UserDeviceOut struct {
	UserDevice
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}

type UserDeviceQuery struct {
	Id            int    `json:"id" gorm:"primaryKey;column:id"`
	AppType       int    `json:"appType" form:"appType,default=-1"`
	DeviceType    int    `json:"deviceType" form:"deviceType,default=-1"`
	DeviceId      string `json:"deviceId" form:"deviceId"`
	DeviceToken   string `json:"deviceToken" form:"deviceToken"`
	UserId        int    `json:"userId" form:"userId,default=-1"`
	AppVersionNum int    `json:"appVersionNum" form:"appVersionNum,default=-1"`
	Status        int    `json:"status" form:"status,default=-1"`
	Start         int    `json:"start" form:"start,default=-1"`
	Size          int    `json:"size" form:"size,default=-1"`
}
