package routers

import (
	"swift_typing_api/app/api"
	"swift_typing_api/app/repos"
	"swift_typing_api/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(

		appAPI *api.AppApi,
		tAllAPI *api.TAllApi,
		userAPI *api.UserApi,
		userDeviceAPI *api.UserDeviceApi,
		authRepo repos.IAuthRepo,
	) error {

		authPath := r.Group("/auth")
		{
			//authPath.Use(middle.ValidateUserToken(authRepo)) //validToken middleware
			authPath.GET("/user", userAPI.GetUserInfo)
			authPath.PUT("/user/:userId", userAPI.UpdateUser)
			authPath.PUT("/user/:userId/password", userAPI.ChangPassword)
			authPath.PUT("/user/:userId/token", userAPI.ChangeToken)
			authPath.POST("/logout", userAPI.Logout)
			/* authPath.GET("/users/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			}) */
		}
		openPath := r.Group("/open")
		{
			openPath.POST("/register", userAPI.AddUser)
			openPath.POST("/userDevice", userDeviceAPI.AddUserDevice)
			openPath.POST("/login", userAPI.Login)
			openPath.GET("/app", appAPI.GetAppInfo)
			openPath.GET("/poem", tAllAPI.GetTPoemInfo)
			openPath.GET("/word", tAllAPI.GetTWordInfo)
			/* openPath.GET("/users/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			}) */
		}
		adminPath := r.Group("/admin")
		{
			adminPath.POST("/app", appAPI.AddApp)
			adminPath.POST("/poem", tAllAPI.AddTPoem)
			adminPath.POST("/word", tAllAPI.AddTWord)
			adminPath.PUT("/app/:appId", appAPI.UpdateApp)
			adminPath.GET("/userDevice", userDeviceAPI.GetUserDevice)
		}
		return nil
	})

	if err != nil {
		common.GetLogger().Error(err)
	}

	return err
}
