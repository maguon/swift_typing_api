package repos

import (
	"swift_typing_api/app/dbs"
	"swift_typing_api/app/models"
)

type ITScoreRepo interface {
	SaveTScore(tScoreInfo *models.TScoreInfo) (int, error)
	GetTScore(tScoredQuery *models.TScoreQuery) (*[]models.TScoreInfoOut, error)
	UpdateTScore(scoreInfo map[string]interface{}) (*models.TScoreInfoOut, error)
}

type TScoreRepo struct {
	db dbs.IDatabase
}

func NewTScoreRepo(db dbs.IDatabase) ITScoreRepo {
	return &TScoreRepo{db: db}
}

func (tScoreRepo *TScoreRepo) SaveTScore(tScoreInfo *models.TScoreInfo) (int, error) {
	result := tScoreRepo.db.GetInstance().Table("t_score").Create(&tScoreInfo)
	return tScoreInfo.Id, result.Error
}

func (tScoreRepo *TScoreRepo) UpdateTScore(scoreInfo map[string]interface{}) (*models.TScoreInfoOut, error) {
	query := "update t_score set score = ? where id = ? RETURNING id "
	queryParamArray := []interface{}{}
	queryParamArray = append(queryParamArray, scoreInfo["score"])
	queryParamArray = append(queryParamArray, scoreInfo["id"])
	var result *models.TScoreInfoOut
	tScoreRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&result)
	return result, nil
}

func (tScoreRepo *TScoreRepo) GetTScore(tScoredQuery *models.TScoreQuery) (*[]models.TScoreInfoOut, error) {
	var resultList *[]models.TScoreInfoOut
	query := "select * from t_score where id is not null "
	queryParamArray := []interface{}{}
	if tScoredQuery.Id > 0 {
		query += " and id = ? "
		queryParamArray = append(queryParamArray, tScoredQuery.Id)
	}
	if tScoredQuery.UserId > 0 {
		query += " and user_id = ? "
		queryParamArray = append(queryParamArray, tScoredQuery.UserId)
	}
	if tScoredQuery.ScoreType > 0 {
		query += " and score_type = ? "
		queryParamArray = append(queryParamArray, tScoredQuery.ScoreType)
	}
	if tScoredQuery.Level >= 0 {
		query += "and level = ? "
		queryParamArray = append(queryParamArray, tScoredQuery.Level)
	}
	if tScoredQuery.Status >= 0 {
		query += "and status = ? "
		queryParamArray = append(queryParamArray, tScoredQuery.Status)
	}

	query += " order by id desc "

	if tScoredQuery.Start >= 0 && tScoredQuery.Size > 0 {
		query += "offset ?  limit ? "
		queryParamArray = append(queryParamArray, tScoredQuery.Start, tScoredQuery.Size)
	}
	tScoreRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&resultList)
	return resultList, nil
}

func (tScoreRepo *TScoreRepo) GetTScoreRank(tScoredQuery *models.TScoreQuery) (*[]models.TScoreInfoOut, error) {
	var resultList *[]models.TScoreInfoOut
	query := "select * from   (select  RANK() OVER (order by score desc) " +
		" row_id,id,user_id,score,correct,incorrect from t_score where score_type = ? and level = ? ) " +
		" rand_score where user_id = ? "
	queryParamArray := []interface{}{}
	queryParamArray = append(queryParamArray, tScoredQuery.ScoreType)
	queryParamArray = append(queryParamArray, tScoredQuery.Level)
	queryParamArray = append(queryParamArray, tScoredQuery.UserId)
	tScoreRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&resultList)
	return resultList, nil
}
