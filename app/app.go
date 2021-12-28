package app

import (
	"swift_typing_api/app/api"
	"swift_typing_api/app/dbs"
	"swift_typing_api/app/repos"
	"swift_typing_api/app/routers"
	"swift_typing_api/common"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	err := dbs.Inject(container)
	if err != nil {
		common.GetLogger().Error("Failed to inject database", err)
	}
	err = repos.Inject(container)
	if err != nil {
		common.GetLogger().Error("Failed to inject repos", err)
	}
	err = api.Inject(container)
	if err != nil {
		common.GetLogger().Error("Failed to inject api", err)
	}
	return container
}

func InitGinEngine(container *dig.Container) *gin.Engine {
	app := gin.New()
	gin.ForceConsoleColor()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Static("/assets", "./assets")
	app.StaticFile("/favicon.ico", "./assets/favicon.ico")
	app.Static("/docs", "./docs")
	app.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/docs/swagger.json"), ginSwagger.InstanceName("Typing Api")))
	err := routers.RegisterAPI(app, container)
	if err != nil {
		common.GetLogger().Error(err)
	}
	return app
}
