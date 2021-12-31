package models

import "time"

type TWordInfo struct {
	Id      int    `json:"id" gorm:"primaryKey;column:id"`
	Word    string `json:"word" `
	Spell   string `json:"spell" `
	Explain string `json:"explain" `
	Example string `json:"example" `
	Refere  string `json:"refere" `
	Level   int    `json:"level" `
	Status  int    `json:"status" `
}
type TWordInfoOut struct {
	TWordInfo
	CreatedOn time.Time `json:"createdOn" `
	UpdatedOn time.Time `json:"updatedOn" `
}
type TWordQuery struct {
	Id     int `json:"id" gorm:"primaryKey;column:id"`
	Level  int `json:"level" form:"level,default=-1"`
	Status int `json:"status" form:"status,default=-1"`
	Start  int `json:"start" form:"start,default=-1"`
	Size   int `json:"size" form:"size,default=-1"`
}
