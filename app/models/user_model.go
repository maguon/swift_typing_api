package models

import "time"

type UserInfo struct {
	UserId   int    `json:"id" gorm:"primaryKey;column:id"`
	Username string `json:"userName" gorm:"column:user_name"`
	RealName string `json:"realName" gorm:"column:real_name"`
	Password string `json:"password" `
	Phone    string `json:"phone" `
	Gender   int    `json:"gender" `
	Type     int    `json:"type" `
	Status   int    `json:"status"`
	Captcha  string `json:"captcha" `
}

type UserInfoOut struct {
	UserInfo
	CreatedOn time.Time `json:"createdOn" `
	UpdatedOn time.Time `json:"updatedOn" `
}
type UserQuery struct {
	UserId   int    `json:"id" form:"userId,default=-1"`
	Username string `json:"userName" form:"userName"`
	RealName string `json:"realName" form:"realName"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Gender   int    `json:"gender" form:"gender,default=-1"`
	Type     int    `json:"type" form:"userType,default=-1"`
	Status   int    `json:"status" form:"status,default=-1"`
	Start    int    `json:"start" form:"start,default=-1"`
	Size     int    `json:"size" form:"size,default=-1"`
}
type UserToken struct {
	UserId      int    `json:"userId" `
	UserType    int    `json:"userType" `
	AccessToken string `json:"accessToken" `
}

type Login struct {
	Username string `json:"username" `
	Password string `json:"password" `
}

type UserPassword struct {
	Password    string `json:"password" `
	NewPassword string `json:"newPassword" `
}
