package repos

import (
	"swift_typing_api/app/dbs"
	"swift_typing_api/app/models"
)

type ITWordRepo interface {
	AddTWord(twordInfo *models.TWordInfo) (int, error)
	GetTWord(tworddQuery *models.TWordQuery) (*[]models.TWordInfoOut, error)
}

type TWordRepo struct {
	db dbs.IDatabase
}

func NewTWordRepo(db dbs.IDatabase) ITWordRepo {
	return &TWordRepo{db: db}
}

func (twordRepo *TWordRepo) AddTWord(twordInfo *models.TWordInfo) (int, error) {
	result := twordRepo.db.GetInstance().Table("t_word").Create(&twordInfo)
	return twordInfo.Id, result.Error
}

func (twordRepo *TWordRepo) GetTWord(tworddQuery *models.TWordQuery) (*[]models.TWordInfoOut, error) {
	var resultList *[]models.TWordInfoOut
	query := "select * from t_word where id is not null "
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
