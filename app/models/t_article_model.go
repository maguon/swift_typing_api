package models

import "time"

type TArticleInfo struct {
	Id      int    `json:"id"  gorm:"primaryKey;column:id"`
	Title   string `json:"title" `
	Author  string `json:"author" `
	Content string `json:"content" `
	Refer   string `json:"refer" `
	Level   int    `json:"level" `
	Status  int    `json:"status" `
}
type TArticleInfoOut struct {
	TArticleInfo
	CreatedOn time.Time `json:"createdOn" `
	UpdatedOn time.Time `json:"updatedOn" `
}
type TArticleQuery struct {
	Id     int `json:"id" gorm:"primaryKey;column:id"`
	Level  int `json:"level" form:"level,default=-1"`
	Status int `json:"status" form:"status,default=-1"`
	Start  int `json:"start" form:"start,default=-1"`
	Size   int `json:"size" form:"size,default=-1"`
}
