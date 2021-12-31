package models

type AuthInfo struct {
	UserId   int `json:"userId" `
	UserType int `json:"userType" `
	Status   int `json:"status" `
}
