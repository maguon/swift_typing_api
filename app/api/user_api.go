package api

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"swift_typing_api/app/models"
	"swift_typing_api/app/repos"
	"swift_typing_api/common"
	"swift_typing_api/util"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	repo repos.IUserRepo
	auth repos.IAuthRepo
}

func NewUserAPI(repo repos.IUserRepo, auth repos.IAuthRepo) *UserApi {
	return &UserApi{repo: repo, auth: auth}
}

// @BasePath /open
// @Summary  user register
// @Schemes
// @Description AddUser
// @Tags user
// @Accept json
// @Param userinfo body models.UserInfo true  "user info "
// @Produce json
// @Success 200 {json} models.AccessToken
// @Router /open/register [post]
func (userApi *UserApi) AddUser(c *gin.Context) {
	var userInfo models.UserInfo
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		common.GetLogger().Error(err)
		util.InvalidParamsReponse(c)
		return
	}
	passwordStr := fmt.Sprintf("%x", md5.Sum([]byte(userInfo.Password)))
	userInfo.Password = strings.ToUpper(passwordStr)
	userId, err := userApi.repo.AddUser(&userInfo)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}
	accessToken, _ := util.GenerateAccess(userId)
	userToken := models.UserToken{
		UserId:      userId,
		UserType:    userInfo.Type,
		AccessToken: accessToken,
	}
	util.SuccessResponse(c, userToken)
}

// @BasePath /open
// @Summary  user login
// @Schemes
// @Description userLogin
// @Tags user
// @Accept json
// @Param userinfo body models.Login true  "user info "
// @Produce json
// @Success 200 {json} models.AccessToken
// @Router /open/login [post]
func (userApi *UserApi) Login(c *gin.Context) {
	var loginInfo models.Login
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		common.GetLogger().Error(err)
		util.InvalidParamsReponse(c)
		return
	}
	paramsMap := make(map[string]interface{})
	paramsMap["phone"] = loginInfo.Username
	userInfoArray, err := userApi.repo.GetUserFullInfo(paramsMap)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}
	if userInfoArray == nil || len(*userInfoArray) == 0 {
		common.GetLogger().Warn("user not exist")
		util.FailedResponse(c, "user not register", nil)
	} else {
		loginPassword := fmt.Sprintf("%x", md5.Sum([]byte(loginInfo.Password)))
		loginPassword = strings.ToUpper(loginPassword)
		if loginPassword == (*userInfoArray)[0].Password {
			accessToken, _ := util.GenerateAccess((*userInfoArray)[0].UserId)
			userToken := models.UserToken{
				UserId:      (*userInfoArray)[0].UserId,
				UserType:    (*userInfoArray)[0].Type,
				AccessToken: accessToken,
			}
			authInfo := models.AuthInfo{
				UserId:   (*userInfoArray)[0].UserId,
				UserType: (*userInfoArray)[0].Type,
				Status:   (*userInfoArray)[0].Status,
			}
			//remove access token from redis
			keys, _ := userApi.auth.GetKeys(strconv.Itoa((*userInfoArray)[0].UserId), "")
			userApi.auth.Remove(keys...)
			err := userApi.auth.Set(accessToken, &authInfo)
			if err != nil {
				common.GetLogger().Error(err)
			}
			common.GetLogger().Info("login success", (*userInfoArray)[0].Phone)
			util.SuccessResponse(c, userToken)
		} else {
			common.GetLogger().Warn("login failed by invalid password %s", (*userInfoArray)[0].Phone)
			util.FailedResponse(c, "login failed by invalid password ", nil)
		}

	}

}

// @BasePath /auth
// @Summary  user logout
// @Schemes
// @Description userLogout
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {json} models.AccessToken
// @Security ApiKeyAuth
// @Router /auth/logout [post]
func (userApi *UserApi) Logout(c *gin.Context) {
	accessToken := c.Request.Header["Auth-Token"]
	if len(accessToken[0]) < 150 {
		common.GetLogger().Error("accessToken Error")
		util.InvalidParamsReponse(c)
		return
	}
	keys, err := userApi.auth.GetKeys("", accessToken[0])
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}
	userApi.auth.Remove(keys...)
	util.SuccessUpdateResponse(c, len(keys))
}

// @BasePath /auth
// @Summary Get User List
// @Schemes
// @Description GetUserInfo
// @Tags user
// @Accept json
// @Param userId query int false "User ID"
// @Param gender query int false "user gender"
// @Param userType query int false "userType"
// @Param status query int false "user status"
// @Param phone query string false "user register phone"
// @Param start query int false "Offset"
// @Param size query int false "Limit"
// @Produce json
// @Success 200 {string} models.AppInfo
// @Router /auth/user [get]
func (userApi *UserApi) GetUserInfo(c *gin.Context) {
	var userQuery models.UserQuery

	c.ShouldBindQuery(&userQuery)
	common.GetLogger().Info(userQuery)
	userOutList, err := userApi.repo.GetUser(&userQuery)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}
	util.SuccessResponse(c, userOutList)
}

// @BasePath /auth
// @Summary  user register
// @Schemes
// @Description AddUser
// @Tags user
// @Accept json
// @Param userId path int true "user ID"
// @Param userinfo body models.UserInfo true  "user info "
// @Produce json
// @Success 200 {json} int
// @Router /auth/user/{userId} [put]
func (userApi *UserApi) UpdateUser(c *gin.Context) {
	var userInfo models.UserInfo
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		common.GetLogger().Error(err)
		util.InvalidParamsReponse(c)
		return
	}
	userInfo.UserId, _ = strconv.Atoi(c.Param("userId"))
	rowsAffected, err := userApi.repo.Update(&userInfo)
	if err != nil {
		common.GetLogger().Error(err)
		util.InternalServerResponse(c)
		return
	}
	util.SuccessUpdateResponse(c, rowsAffected)
}

// @BasePath /auth
// @Summary  user change access token
// @Schemes
// @Description ChangeToken
// @Tags user
// @Accept json
// @Produce json
// @Param userId path int true "user ID"
// @Success 200 {json} models.UserToken
// @Security ApiKeyAuth
// @Router /auth/user/{userId}/token [put]
func (userApi *UserApi) ChangeToken(c *gin.Context) {
	userId := c.Param("userId")
	accessToken := c.Request.Header["Auth-Token"]

	tokenInfo, err := util.ParseToken(accessToken[0])
	if err != nil {
		util.ErrorAuthTokenReponse(c)
	} else {
		if tokenInfo.Id == userId {

			userIdInt, _ := strconv.Atoi(userId)
			userQuery := models.UserQuery{UserId: userIdInt, Gender: -1, Status: -1, Type: -1}
			userInfoList, err := userApi.repo.GetUser(&userQuery)
			if err != nil {
				common.GetLogger().Error(err)
				util.InternalServerResponse(c)
				return
			}
			fmt.Println(userInfoList)
			if userInfoList != nil && len(*userInfoList) > 0 && (*userInfoList)[0].Status == 1 {
				newTokenString, _ := util.GenerateAccess(userIdInt)
				keys, _ := userApi.auth.GetKeys(userId, "")
				userApi.auth.Remove(keys...)
				userApi.auth.Set(newTokenString, &models.AuthInfo{
					UserId:   (*userInfoList)[0].UserId,
					UserType: (*userInfoList)[0].Type,
					Status:   (*userInfoList)[0].Status,
				})
				util.SuccessResponse(c, newTokenString)
			} else {
				util.ErrorAuthTokenReponse(c)
			}
		} else {
			util.ErrorAuthTokenReponse(c)
		}

	}

}