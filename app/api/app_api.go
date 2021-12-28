package api

import (
	"strconv"
	"swift_typing_api/app/models"
	"swift_typing_api/app/repos"
	"swift_typing_api/common"
	"swift_typing_api/util"

	"github.com/gin-gonic/gin"
)

type AppApi struct {
	repo repos.IAppRepo
}

func NewAppAPI(repo repos.IAppRepo) *AppApi {
	return &AppApi{repo: repo}
}

// @BasePath /open
// @Summary Get App List
// @Schemes
// @Description GetAppInfo
// @Tags app
// @Accept json
// @Param appId query int false "App ID"
// @Param appType query int false "App Type"
// @Param deviceType query int false "Device Type"
// @Param forceUpdate query int false "1 force update 0 not"
// @Param start query int false "Offset"
// @Param size query int false "Limit"
// @Produce json
// @Success 200 {string} models.AppInfo
// @Router /open/app [get]
func (appApi *AppApi) GetAppInfo(c *gin.Context) {
	var appQuery models.AppQuery

	c.ShouldBindQuery(&appQuery)
	appList, err := appApi.repo.GetApp(&appQuery)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}
	util.SuccessResponse(c, appList)
}

// @BasePath /admin
// HelloExample godoc
// @Summary add app info
// @Schemes
// @Description AddApp
// @Tags app
// @Accept json
// @Param appinfo body models.AppInfo true  "app info "
// @Produce json
// @Success 200 {json} int
// @Router /admin/app [post]
func (appApi *AppApi) AddApp(c *gin.Context) {
	var appInfo models.AppInfo
	if err := c.ShouldBindJSON(&appInfo); err != nil {
		common.GetLogger().Error(err)
		util.InvalidParamsReponse(c)
		return
	}
	appId, err := appApi.repo.AddApp(&appInfo)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}

	util.SuccessResponse(c, appId)
}

// @BasePath /admin
// HelloExample godoc
// @Summary update app info
// @Schemes
// @Description UpdateApp
// @Tags app
// @Accept json
// @Param appId path int true "App ID"
// @Param appinfo body models.AppInfo true  "app info "
// @Produce json
// @Success 200 {json} int
// @Router /admin/app/{appId} [put]
func (appApi *AppApi) UpdateApp(c *gin.Context) {
	var appInfo models.AppInfo
	err := c.ShouldBindJSON(&appInfo)
	if err != nil {
		common.GetLogger().Error(err.Error())
		util.InvalidParamsReponse(c)
		return
	}
	appInfo.Id, _ = strconv.Atoi(c.Param("appId"))
	common.GetLogger().Debug(appInfo)
	appId, err := appApi.repo.UpdateApp(&appInfo)
	if err != nil {
		common.GetLogger().Error(err.Error())
		util.InvalidParamsReponse(c)
		return
	}
	result := gin.H{
		"message": appId,
	}
	util.SuccessResponse(c, result)
}
