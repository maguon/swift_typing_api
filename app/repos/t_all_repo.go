package repos

import (
	"swift_typing_api/app/dbs"
	"swift_typing_api/app/models"
)

type ITAllRepo interface {
	AddTSentence(tInfo *models.TSentenceInfo) (int, error)
	GetTSentence(tQuery *models.TSentenceQuery) (*[]models.TSentenceInfoOut, error)
	AddTArticle(tInfo *models.TArticleInfo) (int, error)
	GetTArticle(tQuery *models.TArticleQuery) (*[]models.TArticleInfoOut, error)
	AddTPoem(twordInfo *models.TPoemInfo) (int, error)
	GetTPoem(twordQuery *models.TPoemQuery) (*[]models.TPoemInfoOut, error)
	AddTWord(twordInfo *models.TWordInfo) (int, error)
	GetTWord(tworddQuery *models.TWordQuery) (*[]models.TWordInfoOut, error)
	SaveTScore(tScoreInfo *models.TScoreInfo) (int, error)
	UpdateTScore(scoreInfo map[string]interface{}) (*models.TScoreInfoOut, error)
	GetTScoreRank(tScoredQuery *models.TScoreQuery) (*[]models.TScoreInfoOut, error)
	GetTScore(tScoredQuery *models.TScoreQuery) (*[]models.TScoreInfoOut, error)
}

type TAllRepo struct {
	db dbs.IDatabase
}

func NewTAllRepo(db dbs.IDatabase) ITAllRepo {
	return &TAllRepo{db: db}
}

func (tRepo *TAllRepo) AddTSentence(tInfo *models.TSentenceInfo) (int, error) {
	result := tRepo.db.GetInstance().Table("t_sentence").Create(&tInfo)
	return tInfo.Id, result.Error
}

func (tRepo *TAllRepo) GetTSentence(tQuery *models.TSentenceQuery) (*[]models.TSentenceInfoOut, error) {
	var resultList *[]models.TSentenceInfoOut
	query := "select * from t_sentence where id is not null "
	queryParamArray := []interface{}{}
	if tQuery.Id > 0 {
		query += " and id = ? "
		queryParamArray = append(queryParamArray, tQuery.Id)
	}

	if tQuery.Status >= 0 {
		query += "and status = ? "
		queryParamArray = append(queryParamArray, tQuery.Status)
	}

	query += " order by random() "

	if tQuery.Start >= 0 && tQuery.Size > 0 {
		query += "offset ?  limit ? "
		queryParamArray = append(queryParamArray, tQuery.Start, tQuery.Size)
	}
	tRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&resultList)
	return resultList, nil
}

func (tRepo *TAllRepo) AddTArticle(tInfo *models.TArticleInfo) (int, error) {
	result := tRepo.db.GetInstance().Table("t_article").Create(&tInfo)
	return tInfo.Id, result.Error
}

func (tRepo *TAllRepo) GetTArticle(tQuery *models.TArticleQuery) (*[]models.TArticleInfoOut, error) {
	var resultList *[]models.TArticleInfoOut
	query := "select * from t_article where id is not null "
	queryParamArray := []interface{}{}
	if tQuery.Id > 0 {
		query += " and id = ? "
		queryParamArray = append(queryParamArray, tQuery.Id)
	}
	if tQuery.Level >= 0 {
		query += "and level = ? "
		queryParamArray = append(queryParamArray, tQuery.Level)
	}
	if tQuery.Status >= 0 {
		query += "and status = ? "
		queryParamArray = append(queryParamArray, tQuery.Status)
	}

	query += " order by random() "

	if tQuery.Start >= 0 && tQuery.Size > 0 {
		query += "offset ?  limit ? "
		queryParamArray = append(queryParamArray, tQuery.Start, tQuery.Size)
	}
	tRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&resultList)
	return resultList, nil
}

func (tRepo *TAllRepo) AddTPoem(twordInfo *models.TPoemInfo) (int, error) {
	result := tRepo.db.GetInstance().Table("t_word").Create(&twordInfo)
	return twordInfo.Id, result.Error
}

func (tRepo *TAllRepo) GetTPoem(twordQuery *models.TPoemQuery) (*[]models.TPoemInfoOut, error) {
	var resultList *[]models.TPoemInfoOut
	query := "select * from t_poem where id is not null "
	queryParamArray := []interface{}{}
	if twordQuery.Id > 0 {
		query += " and id = ? "
		queryParamArray = append(queryParamArray, twordQuery.Id)
	}
	if twordQuery.Level >= 0 {
		query += "and level = ? "
		queryParamArray = append(queryParamArray, twordQuery.Level)
	}
	if twordQuery.Status >= 0 {
		query += "and status = ? "
		queryParamArray = append(queryParamArray, twordQuery.Status)
	}

	query += " order by random() "

	if twordQuery.Start >= 0 && twordQuery.Size > 0 {
		query += "offset ?  limit ? "
		queryParamArray = append(queryParamArray, twordQuery.Start, twordQuery.Size)
	}
	tRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&resultList)
	return resultList, nil
}

func (tRepo *TAllRepo) AddTWord(twordInfo *models.TWordInfo) (int, error) {
	result := tRepo.db.GetInstance().Table("t_word").Create(&twordInfo)
	return twordInfo.Id, result.Error
}

func (tRepo *TAllRepo) GetTWord(tworddQuery *models.TWordQuery) (*[]models.TWordInfoOut, error) {
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

	query += " order by random() "

	if tworddQuery.Start >= 0 && tworddQuery.Size > 0 {
		query += "offset ?  limit ? "
		queryParamArray = append(queryParamArray, tworddQuery.Start, tworddQuery.Size)
	}
	tRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&resultList)
	return resultList, nil
}

func (tRepo *TAllRepo) SaveTScore(tScoreInfo *models.TScoreInfo) (int, error) {
	result := tRepo.db.GetInstance().Table("t_score").Create(&tScoreInfo)
	return tScoreInfo.Id, result.Error
}

func (tRepo *TAllRepo) UpdateTScore(scoreInfo map[string]interface{}) (*models.TScoreInfoOut, error) {
	query := "update t_score set score = ? where id = ? RETURNING id "
	queryParamArray := []interface{}{}
	queryParamArray = append(queryParamArray, scoreInfo["score"])
	queryParamArray = append(queryParamArray, scoreInfo["id"])
	var result *models.TScoreInfoOut
	tRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&result)
	return result, nil
}

func (tRepo *TAllRepo) GetTScore(tScoredQuery *models.TScoreQuery) (*[]models.TScoreInfoOut, error) {
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
	tRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&resultList)
	return resultList, nil
}

func (tRepo *TAllRepo) GetTScoreRank(tScoredQuery *models.TScoreQuery) (*[]models.TScoreInfoOut, error) {
	var resultList *[]models.TScoreInfoOut
	query := "select * from   (select  RANK() OVER (order by score desc) " +
		" row_id,id,user_id,score,correct,incorrect from t_score where score_type = ? and level = ? ) " +
		" rand_score where user_id = ? "
	queryParamArray := []interface{}{}
	queryParamArray = append(queryParamArray, tScoredQuery.ScoreType)
	queryParamArray = append(queryParamArray, tScoredQuery.Level)
	queryParamArray = append(queryParamArray, tScoredQuery.UserId)
	tRepo.db.GetInstance().Raw(query, queryParamArray...).Scan(&resultList)
	return resultList, nil
}
