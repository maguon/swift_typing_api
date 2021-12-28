package repos

import (
	"swift_typing_api/app/dbs"
	"swift_typing_api/app/models"
)

type IAppRepo interface {
	AddApp(appInfo *models.AppInfo) (int, error)
	GetApp(appQuery *models.AppQuery) (*[]models.AppInfoOut, error)
	UpdateApp(appInfo *models.AppInfo) (int, error)
}

type AppRepo struct {
	db dbs.IDatabase
}

func NewAppRepo(db dbs.IDatabase) IAppRepo {
	return &AppRepo{db: db}
}

func (appRepo *AppRepo) AddApp(appInfo *models.AppInfo) (int, error) {
	result := appRepo.db.GetInstance().Table("app_info").Create(&appInfo)
	return appInfo.Id, result.Error
}

func (appRepo *AppRepo) UpdateApp(appInfo *models.AppInfo) (int, error) {
	result := appRepo.db.GetInstance().Table("app_info").Updates(&appInfo)
	return int(result.RowsAffected), result.Error
}

func (appRepo *AppRepo) GetApp(appQuery *models.AppQuery) (*[]models.AppInfoOut, error) {
	var appInfoList *[]models.AppInfoOut
	query := "select * from app_info where id is not null "
	queryParamArray := []interface{}{}
	if appQuery.Id > 0 {
		query += " and id = ? "
		queryParamArray = append(queryParamArray, appQuery.Id)
	}
	if appQuery.AppType >= 0 {
		query += "and app_type = ? "
		queryParamArray = append(queryParamArray, appQuery.AppType)
	}
	if appQuery.DeviceType >= 0 {
		query += "and device_type = ? "
		queryParamArray = append(queryParamArray, appQuery.DeviceType)
	}
	if appQuery.ForceUpdate >= 0 {
		query += "and force_update = ? "
		queryParamArray = append(queryParamArray, appQuery.ForceUpdate)
	}
	if appQuery.Status >= 0 {
		query += "and status = ? "
		queryParamArray = append(queryParamArray, appQuery.Status)
	}

	query += " order by id desc "

	if appQuery.Start >= 0 && appQuery.Size > 0 {
		query += "offset ?  limit ? "
		queryParamArray = append(queryParamArray, appQuery.Start, appQuery.Size)
	}
	appRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&appInfoList)
	return appInfoList, nil
}
