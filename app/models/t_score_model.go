package models

import "time"

type TScoreInfo struct {
	Id        int     `json:"id" gorm:"primaryKey;column:id"`
	UserId    int     `json:"userId" `
	Correct   int     `json:"correct" `
	Incorrect int     `json:"incorrect" `
	Score     float32 `json:"score" `
	ScoreType int     `json:"scoreType" `
	Level     int     `json:"level" `
	Status    int     `json:"status" `
}
type TScoreInfoOut struct {
	TScoreInfo
	RowId     int       `json:"rowId" `
	CreatedOn time.Time `json:"createdOn" `
	UpdatedOn time.Time `json:"updatedOn" `
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
