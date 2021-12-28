package repos

import (
	"fmt"
	"swift_typing_api/app/dbs"
	"swift_typing_api/app/models"
)

type IUserRepo interface {
	AddUser(userInfo *models.UserInfo) (int, error)
	GetUser(userQuery *models.UserQuery) (*[]models.UserInfoOut, error)
	GetUserFullInfo(userInfo map[string]interface{}) (*[]models.UserInfo, error)
	Update(userInfo *models.UserInfo) (int, error)
	//Login(login *models.Login) (*models.UserToken, error)
}
type UserRepo struct {
	db dbs.IDatabase
}

// NewUserRepository return new IUserRepository interface
func NewUserRepo(db dbs.IDatabase) IUserRepo {
	return &UserRepo{db: db}
}

func (userRepo *UserRepo) AddUser(userInfo *models.UserInfo) (int, error) {
	result := userRepo.db.GetInstance().Table("user_info").Create(&userInfo)
	return userInfo.UserId, result.Error
}

func (userRepo *UserRepo) GetUser(userQuery *models.UserQuery) (*[]models.UserInfoOut, error) {
	var userInfoList *[]models.UserInfoOut
	query := "select id,user_name,real_name,status,phone,email,gender,type,created_on,updated_on from user_info where id is not null "
	queryParamArray := []interface{}{}
	if userQuery.UserId > 0 {
		query += " and id = ? "
		queryParamArray = append(queryParamArray, userQuery.UserId)
	}
	if userQuery.Gender >= 0 {
		query += "and gender = ? "
		queryParamArray = append(queryParamArray, userQuery.Gender)
	}
	if userQuery.Status >= 0 {
		query += "and status = ? "
		queryParamArray = append(queryParamArray, userQuery.Gender)
	}
	if userQuery.Type >= 0 {
		query += "and type = ? "
		queryParamArray = append(queryParamArray, userQuery.Type)
	}
	if userQuery.Phone != "" {
		query += "and phone = ? "
		queryParamArray = append(queryParamArray, userQuery.Phone)
	}
	if userQuery.Email != "" {
		query += "and email = ? "
		queryParamArray = append(queryParamArray, userQuery.Email)
	}
	query += " order by id desc "

	if userQuery.Start >= 0 && userQuery.Size > 0 {
		query += "offset ?  limit ? "
		queryParamArray = append(queryParamArray, userQuery.Start, userQuery.Size)
	}
	fmt.Println(query)
	userRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&userInfoList)
	return userInfoList, nil
}

func (userRepo *UserRepo) GetUserFullInfo(userInfoQuery map[string]interface{}) (*[]models.UserInfo, error) {
	var userInfoList *[]models.UserInfo
	query := "select * from user_info where id is not null "
	queryParamArray := []interface{}{}
	if userInfoQuery["userId"] != nil {
		query += " and id = ? "
		queryParamArray = append(queryParamArray, userInfoQuery["userId"])
	}
	if userInfoQuery["gender"] != nil {
		query += "and gender = ? "
		queryParamArray = append(queryParamArray, userInfoQuery["userId"])
	}
	if userInfoQuery["status"] != nil {
		query += "and status = ? "
		queryParamArray = append(queryParamArray, userInfoQuery["status"])
	}
	if userInfoQuery["userType"] != nil {
		query += "and type = ? "
		queryParamArray = append(queryParamArray, userInfoQuery["userType"])
	}
	if userInfoQuery["phone"] != nil {
		query += "and phone = ? "
		queryParamArray = append(queryParamArray, userInfoQuery["phone"])
	}
	if userInfoQuery["email"] != nil {
		query += "and email = ? "
		queryParamArray = append(queryParamArray, userInfoQuery["email"])
	}
	query += " order by id desc "

	if userInfoQuery["start"] != nil && userInfoQuery["size"] != nil {
		query += "offset ?  limit ? "
		queryParamArray = append(queryParamArray, userInfoQuery["start"], userInfoQuery["size"])
	}
	userRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&userInfoList)
	return userInfoList, nil
}

func (userRepo *UserRepo) Update(userInfo *models.UserInfo) (int, error) {
	result := userRepo.db.GetInstance().Table("user_info").Updates(&userInfo)
	return int(result.RowsAffected), result.Error
}
