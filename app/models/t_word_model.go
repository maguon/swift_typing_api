package models

import "time"

type TWordInfo struct {
	Id      int `gorm:"primaryKey;column:id"`
	Word    string
	Spell   string
	Explain string
	Example string
	Refere  string
	Level   int
	Status  int
}
type TWordInfoOut struct {
	Id        int `gorm:"primaryKey;column:id"`
	Word      string
	Spell     string
	Explain   string
	Example   string
	Refere    string
	Level     int
	Status    int
	CreatedOn time.Time
	UpdatedOn time.Time
}
type TWordQuery struct {
	Id     int `json:"id" gorm:"primaryKey;column:id"`
	Level  int `json:"level" form:"level,default=-1"`
	Status int `json:"status" form:"status,default=-1"`
	Start  int `json:"start" form:"start,default=-1"`
	Size   int `json:"size" form:"size,default=-1"`
}
