package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"swift_typing_api/app"
	"swift_typing_api/common"
)

// @title Swift API Documents
// @version 1.0
// @description Swift Typing Api Docs.

// @in header
// @name Authorization

// @contact.name myxxjs
// @contact.url http://www.myxxjs.com
// @contact.email info@myxxjs.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name auth-token
func main() {
	container := app.BuildContainer()
	engine := app.InitGinEngine(container)

	server := &http.Server{
		Addr:    ":8888",
		Handler: engine,
	}
	common.GetLogger().Info("Listen at:", server.Addr)
	go func() {

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			common.GetLogger().Fatal(err)

			common.GetLogger().Fatalf("Error: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	common.GetLogger().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		common.GetLogger().Fatal("Server Shutdown: ", err)

	}
	// catching ctx.Done(). timeout of 1 seconds.
	select {
	case <-ctx.Done():
		common.GetLogger().Info("Timeout of 1 seconds.")

	}
	common.GetLogger().Info("Server exiting")
}
