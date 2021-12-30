package models

import "time"

type UserInfo struct {
	UserId   int    `json:"id" gorm:"primaryKey;column:id"`
	Username string ` gorm:"column:user_name"`
	RealName string `gorm:"column:real_name"`
	Password string `json:"password" `
	Phone    string `json:"phone" `
	Gender   int    `json:"gender" `
	Type     int    `json:"type" `
	Status   int    `json:"status" `
}

type UserInfoOut struct {
	UserId    int    `gorm:"column:id"`
	Username  string `gorm:"column:user_name"`
	RealName  string
	Phone     string
	Email     string
	Gender    int
	Type      int
	Status    int
	CreatedOn time.Time
	UpdatedOn time.Time
}
type UserQuery struct {
	UserId   int    `json:"id" form:"userId,default=-1"`
	Username string `json:"user_name" form:"userName"`
	RealName string `json:"real_name" form:"realName"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Gender   int    `json:"gender" form:"gender,default=-1"`
	Type     int    `json:"type" form:"userType,default=-1"`
	Status   int    `json:"status" form:"status,default=-1"`
	Start    int    `json:"start" form:"start,default=-1"`
	Size     int    `json:"size" form:"size,default=-1"`
}
type UserToken struct {
	UserId      int
	UserType    int
	AccessToken string
}

type Login struct {
	Username string `json:"username" `
	Password string `json:"password" `
}

type UserPassword struct {
	Password    string
	NewPassword string
}
