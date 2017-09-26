package main

import (
	"fmt"
	"telehorn/config"

	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/github"
	"github.com/gin-gonic/gin"
)

func init() {
	config.InitLog("api")
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	InitRouting(router)

	// -- Async start bot
	go func() {
		StartBot()
	}()

	// -- And start API
	router.Run(appPort()) // listen and serve on 0.0.0.0:appPort
}

func appPort() string {
	configStruct := config.Get()
	return fmt.Sprintf("0.0.0.0:%s", configStruct.Port)
}
