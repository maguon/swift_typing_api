package repos

import (
	"swift_typing_api/app/dbs"
	"swift_typing_api/app/models"
)

type ITPoemRepo interface {
	AddTPoem(twordInfo *models.TPoemInfo) (int, error)
	GetTPoem(tworddQuery *models.TPoemQuery) (*[]models.TPoemInfoOut, error)
}

type TPoemRepo struct {
	db dbs.IDatabase
}

func NewTPoemRepo(db dbs.IDatabase) ITPoemRepo {
	return &TPoemRepo{db: db}
}

func (twordRepo *TPoemRepo) AddTPoem(twordInfo *models.TPoemInfo) (int, error) {
	result := twordRepo.db.GetInstance().Table("t_word").Create(&twordInfo)
	return twordInfo.Id, result.Error
}

func (twordRepo *TPoemRepo) GetTPoem(tworddQuery *models.TPoemQuery) (*[]models.TPoemInfoOut, error) {
	var resultList *[]models.TPoemInfoOut
	query := "select * from t_poem where id is not null "
	queryParamArray := []interface{}{}
	if tworddQuery.Id > 0 {
		query += " and id = ? "
		queryParamArray = append(queryParamArray, tworddQuery.Id)
	}
	if tworddQuery.Level >= 0 {
		query += "and level = ? "
		queryParamArray = append(queryParamArray, tworddQuery.Level)
	}
	if tworddQuery.Status >= 0 {
		query += "and status = ? "
		queryParamArray = append(queryParamArray, tworddQuery.Status)
	}

	query += " order by id desc "

	if tworddQuery.Start >= 0 && tworddQuery.Size > 0 {
		query += "offset ?  limit ? "
		queryParamArray = append(queryParamArray, tworddQuery.Start, tworddQuery.Size)
	}
	twordRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&resultList)
	return resultList, nil
}
