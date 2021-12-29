package middle

import (
	"swift_typing_api/app/models"
	"swift_typing_api/app/repos"
	"swift_typing_api/util"

	"github.com/gin-gonic/gin"
)

func ValidateUserToken(authRepo repos.IAuthRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("auth-token")
		var authInfo models.AuthInfo
		authRepo.Get(authToken, &authInfo)
		if authInfo.UserId == 0 {
			util.ErrorAuthTokenReponse(c)
			c.Abort()
			return
		} else {
			c.Set("_userId", authInfo.UserId)
			c.Next()
		}

	}
}

func ValidateAdminToken(authRepo repos.IAuthRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
