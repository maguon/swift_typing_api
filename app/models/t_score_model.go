package models

import "time"

type TScoreInfo struct {
	Id        int `gorm:"primaryKey;column:id"`
	UserId    int
	Correct   int
	Incorrect int
	Score     float32
	ScoreType int
	Level     int
	Status    int
}
type TScoreInfoOut struct {
	Id        int `gorm:"primaryKey;column:id"`
	RowId     int
	UserId    int
	Correct   int
	Incorrect int
	Score     float32
	ScoreType int
	Level     int
	Status    int
	CreatedOn time.Time
	UpdatedOn time.Time
}
type TScoreQuery struct {
	Id        int `json:"id" gorm:"primaryKey;column:id"`
	Level     int `json:"level" form:"level,default=-1"`
	Status    int `json:"status" form:"status,default=-1"`
	UserId    int `json:"userId" form:"userId,default=-1"`
	ScoreType int `json:"scoreType" form:"scoreType,default=-1"`
	Start     int `json:"start" form:"start,default=-1"`
	Size      int `json:"size" form:"size,default=-1"`
}
