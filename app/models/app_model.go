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
	Id            int    `json:"id" gorm:"primaryKey;column:id"`
	AppType       int    `json:"appType" `
	DeviceType    int    `json:"deviceType" `
	Version       string `json:"version" `
	VersionNum    int    `json:"versionNum" `
	MinVersionNum int    `json:"minVersionNum" `
	ForceUpdate   int    `json:"forceUpdate" `
	Status        int    `json:"status" `
	Url           string `json:"url" `
	Remarks       string `json:"remark" `
}
type AppInfoOut struct {
	AppInfo
	CreatedOn time.Time `json:"createdOn" `
	UpdatedOn time.Time `json:"updatedOn" `
}

type AppQuery struct {
	Id          int `json:"id" gorm:"primaryKey;column:id"`
	AppType     int `json:"appType" form:"app_type,default=-1"`
	DeviceType  int `json:"deviceType" form:"device_type,default=-1"`
	ForceUpdate int `json:"forceUpdate" form:"force_update,default=-1"`
	Status      int `json:"status" form:"status,default=-1"`
	Start       int `json:"start" form:"start,default=-1"`
	Size        int `json:"size" form:"size,default=-1"`
}
