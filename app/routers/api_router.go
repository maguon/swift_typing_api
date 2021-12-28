package routers

import (
	"swift_typing_api/app/api"
	"swift_typing_api/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		userAPI *api.UserApi,
		appAPI *api.AppApi,
	) error {
		authPath := r.Group("/auth")
		{
			authPath.GET("/user", userAPI.GetUserInfo)
			authPath.PUT("/user/:userId", userAPI.UpdateUser)
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
			openPath.POST("/login", userAPI.Login)
			openPath.GET("/app", appAPI.GetAppInfo)
			/* openPath.GET("/users/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			}) */
		}
		adminPath := r.Group("/admin")
		{
			adminPath.POST("/app", appAPI.AddApp)
			adminPath.PUT("/app/:appId", appAPI.UpdateApp)
		}
		return nil
	})

	if err != nil {
		common.GetLogger().Error(err)
	}

	return err
}