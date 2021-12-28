package api

import (
	"swift_typing_api/app/models"
	"swift_typing_api/app/repos"
	"swift_typing_api/common"
	"swift_typing_api/util"

	"github.com/gin-gonic/gin"
)

type UserDeviceApi struct {
	repo repos.IUserDeviceRepo
}

func NewUserDeviceAPI(repo repos.IUserDeviceRepo) *UserDeviceApi {
	return &UserDeviceApi{repo: repo}
}

// @BasePath /admin
// @Summary Get App List
// @Schemes
// @Description GetAppInfo
// @Tags UserDevice
// @Accept json
// @Param appType query int false "App Type"
// @Param deviceType query int false "Device Type"
// @Param userId query int false "User Id"
// @Param status query int false "status"
// @Param appVersionNum query int false "app version num"
// @Param start query int false "Offset"
// @Param size query int false "Limit"
// @Produce json
// @Success 200 {string} models.AppInfo
// @Router /admin/userDevice [get]
func (userDeviceApi *UserDeviceApi) GetUserDevice(c *gin.Context) {
	var userDeviceQuery models.UserDeviceQuery
	c.ShouldBindQuery(&userDeviceQuery)
	userDeviceList, err := userDeviceApi.repo.GetUserDevice(&userDeviceQuery)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}
	util.SuccessResponse(c, userDeviceList)
}

// @BasePath /open
// HelloExample godoc
// @Summary add app info
// @Schemes
// @Description AddApp
// @Tags UserDevice
// @Accept json
// @Param appinfo body models.UserDevice true  "app info "
// @Produce json
// @Success 200 {json} int
// @Router /open/userDevice [post]
func (userDeviceApi *UserDeviceApi) AddUserDevice(c *gin.Context) {
	var userDevice models.UserDevice
	if err := c.ShouldBindJSON(&userDevice); err != nil {
		common.GetLogger().Error(err)
		util.InvalidParamsReponse(c)
		return
	}
	userDeviceQuery := models.UserDeviceQuery{
		DeviceId: userDevice.DeviceId,
		Status:   1,
	}
	userDeviceList, err := userDeviceApi.repo.GetUserDevice(&userDeviceQuery)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}

	if userDeviceList == nil {
		userDeviceId, err := userDeviceApi.repo.AddUserDevice(&userDevice)
		if err != nil {
			common.GetLogger().Error(err)
			util.InternalServerResponse(c)
			return
		}
		util.SuccessResponse(c, userDeviceId)
	} else if len(*userDeviceList) > 0 && (*userDeviceList)[0].UserId != userDevice.UserId {
		userDeviceId, err := userDeviceApi.repo.AddUserDevice(&userDevice)
		if err != nil {
			common.GetLogger().Error(err)
			util.InternalServerResponse(c)
			return
		}

		util.SuccessResponse(c, userDeviceId)
	} else {
		util.SuccessResponse(c, (*userDeviceList)[0].Id)
	}
}
