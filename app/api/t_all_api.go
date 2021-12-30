package api

import (
	"strconv"
	"swift_typing_api/app/models"
	"swift_typing_api/app/repos"
	"swift_typing_api/common"
	"swift_typing_api/util"

	"github.com/gin-gonic/gin"
)

type TAllApi struct {
	wordRepo  repos.ITWordRepo
	poemRepo  repos.ITPoemRepo
	scoreRepo repos.ITScoreRepo
}

func NewTAllAPI(wordRepo repos.ITWordRepo, poemRepo repos.ITPoemRepo, scoreRepo repos.ITScoreRepo) *TAllApi {
	return &TAllApi{wordRepo: wordRepo, poemRepo: poemRepo, scoreRepo: scoreRepo}
}

// @BasePath /open
// @Summary Get TPoem List
// @Schemes
// @Description GetTPoemInfo
// @Tags T
// @Accept json
// @Param id query int false "Pome ID"
// @Param level query int false "poem level"
// @Param status query int false "pome status"
// @Param start query int false "Offset"
// @Param size query int false "Limit"
// @Produce json
// @Success 200 {string} models.TPomeInfoOut
// @Router /open/poem [get]
func (tAllApi *TAllApi) GetTPoemInfo(c *gin.Context) {
	var tPoemQuery models.TPoemQuery

	c.ShouldBindQuery(&tPoemQuery)
	appList, err := tAllApi.poemRepo.GetTPoem(&tPoemQuery)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}
	util.SuccessResponse(c, appList)
}

// @BasePath /admin
// @Summary add TPome info
// @Schemes
// @Description AddTPoem
// @Tags T
// @Accept json
// @Param appinfo body models.TPoemInfo true  "poem info "
// @Produce json
// @Success 200 {json} int
// @Security ApiKeyAuth
// @Router /admin/poem [post]
func (tAllApi *TAllApi) AddTPoem(c *gin.Context) {
	var tPomeInfo models.TPoemInfo
	if err := c.ShouldBindJSON(&tPomeInfo); err != nil {
		common.GetLogger().Error(err)
		util.InvalidParamsReponse(c)
		return
	}
	tPomeId, err := tAllApi.poemRepo.AddTPoem(&tPomeInfo)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}

	util.SuccessResponse(c, tPomeId)
}

// @BasePath /open
// @Summary Get TWord List
// @Schemes
// @Description GetTWordInfo
// @Tags T
// @Accept json
// @Param id query int false "Word ID"
// @Param level query int false "word level"
// @Param status query int false "word status"
// @Param start query int false "Offset"
// @Param size query int false "Limit"
// @Produce json
// @Success 200 {string} models.TWordInfoOut
// @Router /open/word [get]
func (tAllApi *TAllApi) GetTWordInfo(c *gin.Context) {
	var tWordQuery models.TWordQuery

	c.ShouldBindQuery(&tWordQuery)
	resultList, err := tAllApi.wordRepo.GetTWord(&tWordQuery)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}
	util.SuccessResponse(c, resultList)
}

// @BasePath /admin
// @Summary add TWord info
// @Schemes
// @Description AddTWord
// @Tags T
// @Accept json
// @Param appinfo body models.TWordInfo true  "word info "
// @Produce json
// @Success 200 {json} int
// @Security ApiKeyAuth
// @Router /admin/word [post]
func (tAllApi *TAllApi) AddTWord(c *gin.Context) {
	var tWordInfo models.TWordInfo
	if err := c.ShouldBindJSON(&tWordInfo); err != nil {
		common.GetLogger().Error(err)
		util.InvalidParamsReponse(c)
		return
	}
	tWordId, err := tAllApi.wordRepo.AddTWord(&tWordInfo)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}

	util.SuccessResponse(c, tWordId)
}

// @BasePath /auth
// @Summary save Score info
// @Schemes
// @Description SaveScore
// @Tags T
// @Accept json
// @Param appinfo body models.TScoreInfo true  "score info "
// @Produce json
// @Success 200 {json} int
// @Security ApiKeyAuth
// @Router /auth/score [post]
func (tAllApi *TAllApi) SaveTScore(c *gin.Context) {
	var tScoreInfo models.TScoreInfo
	if err := c.ShouldBindJSON(&tScoreInfo); err != nil {
		common.GetLogger().Error(err)
		util.InvalidParamsReponse(c)
		return
	}
	tScoreQuery := models.TScoreQuery{
		UserId:    tScoreInfo.UserId,
		Level:     tScoreInfo.Level,
		ScoreType: tScoreInfo.ScoreType,
	}
	tScoreList, err := tAllApi.scoreRepo.GetTScore(&tScoreQuery)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}
	if tScoreList == nil || len(*tScoreList) < 1 {
		tScoreId, err := tAllApi.scoreRepo.SaveTScore(&tScoreInfo)
		if err != nil {
			common.GetLogger().Error(err)
			util.InternalServerResponse(c)
			return
		}

		util.SuccessResponse(c, tScoreId)
	} else if (*tScoreList)[0].Score < tScoreInfo.Score {
		paramsMap := make(map[string]interface{})
		paramsMap["id"] = (*tScoreList)[0].Id
		paramsMap["score"] = tScoreInfo.Score
		tAllApi.scoreRepo.UpdateTScore(paramsMap)
	} else {
		util.SuccessResponse(c, (*tScoreList)[0].Id)
	}

}

// @BasePath /open
// @Summary Get TScore List
// @Schemes
// @Description GetTScoreInfo
// @Tags T
// @Accept json
// @Param id query int false "Score ID"
// @Param level query int false "word level"
// @Param userId query int false "user id"
// @Param scoreType query int false "score type"
// @Param status query int false "word status"
// @Param start query int false "Offset"
// @Param size query int false "Limit"
// @Produce json
// @Success 200 {string} models.TScoreInfoOut
// @Router /open/score [get]
func (tAllApi *TAllApi) GetTScoreInfo(c *gin.Context) {
	var tScoreQuery models.TScoreQuery

	c.ShouldBindQuery(&tScoreQuery)
	resultList, err := tAllApi.scoreRepo.GetTScore(&tScoreQuery)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}
	util.SuccessResponse(c, resultList)
}

// @BasePath /auth
// @Summary Get User TScore Rank
// @Schemes
// @Description GetUserScoreRank
// @Tags T
// @Accept json
// @Param userId path int true "user id"
// @Param id query int false "Score ID"
// @Param level query int false "word level"
// @Param scoreType query int false "score type"
// @Param start query int false "Offset"
// @Param size query int false "Limit"
// @Produce json
// @Success 200 {string} models.TScoreInfoOut
// @Router /auth/user/{userId}/scoreRank [get]
func (tAllApi *TAllApi) GetUserScoreRank(c *gin.Context) {
	var tScoreQuery models.TScoreQuery

	c.ShouldBindQuery(&tScoreQuery)
	userId := c.Param("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		common.GetLogger().Error(err)
		util.InvalidParamsReponse(c)
		return
	}
	tScoreQuery.UserId = userIdInt
	resultList, err := tAllApi.scoreRepo.GetTScore(&tScoreQuery)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}
	util.SuccessResponse(c, resultList)
}
