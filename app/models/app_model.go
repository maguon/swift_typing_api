package models

import "time"

/* type AppInfo struct {
	Id            int       `json:"id" gorm:"primaryKey;column:id"`
	AppType       int       `json:"app_type" `
	DeviceType    int       `json:"device_type" `
	Version       string    `json:"version" `
	VersionNum    string    `json:"version_num" `
	MinVersionNum string    `json:"min_version_num" `
	ForceUpdate   int       `json:"force_update" gorm:"default=0"`
	Status        int       `json:"status" gorm:"default=1" `
	Url           string    `json:"url" `
	Remarks       string    `json:"remarks" `
	CreatedOn     time.Time `gorm:"default=time.Now()"`
	UpdatedOn     time.Time `gorm:"default=time.Now()"`
} */
type AppInfo struct {
	Id            int `gorm:"primaryKey;column:id"`
	AppType       int
	DeviceType    int
	Version       string
	VersionNum    int
	MinVersionNum int
	ForceUpdate   int
	Status        int
	Url           string
	Remarks       string
}
type AppInfoOut struct {
	Id            int
	AppType       int
	DeviceType    int
	Version       string
	VersionNum    int
	MinVersionNum int
	ForceUpdate   int
	Status        int
	Url           string
	Remarks       string
	CreatedOn     time.Time
	UpdatedOn     time.Time
}

type AppQuery struct {
	Id          int `json:"id" gorm:"primaryKey;column:id"`
	AppType     int `json:"app_type" form:"app_type,default=-1"`
	DeviceType  int `json:"device_type" form:"device_type,default=-1"`
	ForceUpdate int `json:"force_update" form:"force_update,default=-1"`
	Status      int `json:"status" form:"status,default=-1"`
	Start       int `json:"start" form:"start,default=-1"`
	Size        int `json:"size" form:"size,default=-1"`
}
