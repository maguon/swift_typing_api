package models

import "time"

type TSentenceInfo struct {
	Id      int    `json:"id"  gorm:"primaryKey;column:id"`
	Author  string `json:"author" `
	Content string `json:"content" `
	Status  int    `json:"status" `
}
type TSentenceInfoOut struct {
	TSentenceInfo
	CreatedOn time.Time `json:"createdOn" `
	UpdatedOn time.Time `json:"updatedOn" `
}
type TSentenceQuery struct {
	Id     int `json:"id" gorm:"primaryKey;column:id"`
	Status int `json:"status" form:"status,default=-1"`
	Start  int `json:"start" form:"start,default=-1"`
	Size   int `json:"size" form:"size,default=-1"`
}
