package models

import "time"

type TPoemInfo struct {
	Id      int `gorm:"primaryKey;column:id"`
	Title   string
	Author  string
	Content string
	Level   int
	Status  int
}
type TPoemInfoOut struct {
	Id        int `gorm:"primaryKey;column:id"`
	Title     string
	Author    string
	Content   string
	Level     int
	Status    int
	CreatedOn time.Time
	UpdatedOn time.Time
}
type TPoemQuery struct {
	Id     int `json:"id" gorm:"primaryKey;column:id"`
	Level  int `json:"level" form:"level,default=-1"`
	Status int `json:"status" form:"status,default=-1"`
	Start  int `json:"start" form:"start,default=-1"`
	Size   int `json:"size" form:"size,default=-1"`
}
