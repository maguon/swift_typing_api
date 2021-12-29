package api

import (
	"swift_typing_api/app/models"
	"swift_typing_api/app/repos"
	"swift_typing_api/common"
	"swift_typing_api/util"

	"github.com/gin-gonic/gin"
)

type TAllApi struct {
	wordRepo repos.ITWordRepo
	poemRepo repos.ITPoemRepo
}

func NewTAllAPI(wordRepo repos.ITWordRepo, poemRepo repos.ITPoemRepo) *TAllApi {
	return &TAllApi{wordRepo: wordRepo, poemRepo: poemRepo}
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
