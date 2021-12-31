package repos

import (
	"swift_typing_api/app/dbs"
	"swift_typing_api/app/models"
)

type IUserDeviceRepo interface {
	AddUserDevice(appInfo *models.UserDevice) (int, error)
	GetUserDevice(appQuery *models.UserDeviceQuery) (*[]models.UserDeviceOut, error)
	UpdateUserDeviceStatus(userDeviceQuery *models.UserDeviceQuery) (*[]models.UserDeviceOut, error)
}

type UserDeviceRepo struct {
	db dbs.IDatabase
}

func NewUserDeviceRepo(db dbs.IDatabase) IUserDeviceRepo {
	return &UserDeviceRepo{db: db}
}

func (userDeviceRepo *UserDeviceRepo) AddUserDevice(userDeviceInfo *models.UserDevice) (int, error) {
	result := userDeviceRepo.db.GetInstance().Table("user_device_info").Create(&userDeviceInfo)
	return userDeviceInfo.Id, result.Error
}
func (userDeviceRepo *UserDeviceRepo) UpdateUserDeviceStatus(userDeviceQuery *models.UserDeviceQuery) (*[]models.UserDeviceOut, error) {

	var resultList *[]models.UserDeviceOut
	query := "update user_device_info set status = ?  where id is not null "
	queryParamArray := []interface{}{}
	queryParamArray = append(queryParamArray, userDeviceQuery.Status)
	if userDeviceQuery.Id > 0 {
		query += " and id = ? "
		queryParamArray = append(queryParamArray, userDeviceQuery.Id)
	}
	if userDeviceQuery.UserId > 0 {
		query += " and user_id = ? "
		queryParamArray = append(queryParamArray, userDeviceQuery.UserId)
	}
	if userDeviceQuery.DeviceId != "" {
		query += " and device_id = ? "
		queryParamArray = append(queryParamArray, userDeviceQuery.DeviceId)
	}
	query += " RETURNING id "
	userDeviceRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&resultList)
	return resultList, nil
}

func (userDeviceRepo *UserDeviceRepo) GetUserDevice(userDeviceQuery *models.UserDeviceQuery) (*[]models.UserDeviceOut, error) {
	var resultList *[]models.UserDeviceOut
	query := "select * from user_device_info where id is not null "
	queryParamArray := []interface{}{}
	if userDeviceQuery.Id > 0 {
		query += " and id = ? "
		queryParamArray = append(queryParamArray, userDeviceQuery.Id)
	}
	if userDeviceQuery.AppType > 0 {
		query += "and app_type = ? "
		queryParamArray = append(queryParamArray, userDeviceQuery.AppType)
	}
	if userDeviceQuery.AppVersionNum > 0 {
		query += "and app_version_num = ? "
		queryParamArray = append(queryParamArray, userDeviceQuery.AppVersionNum)
	}
	if userDeviceQuery.DeviceType > 0 {
		query += "and device_type = ? "
		queryParamArray = append(queryParamArray, userDeviceQuery.DeviceType)
	}
	if userDeviceQuery.DeviceId != "" {
		query += "and device_id = ? "
		queryParamArray = append(queryParamArray, userDeviceQuery.DeviceId)
	}
	if userDeviceQuery.DeviceToken != "" {
		query += "and device_id = ? "
		queryParamArray = append(queryParamArray, userDeviceQuery.DeviceToken)
	}
	if userDeviceQuery.Status >= 0 {
		query += "and status = ? "
		queryParamArray = append(queryParamArray, userDeviceQuery.Status)
	}

	query += " order by id desc "

	if userDeviceQuery.Start >= 0 && userDeviceQuery.Size > 0 {
		query += "offset ?  limit ? "
		queryParamArray = append(queryParamArray, userDeviceQuery.Start, userDeviceQuery.Size)
	}
	userDeviceRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&resultList)
	return resultList, nil
}
