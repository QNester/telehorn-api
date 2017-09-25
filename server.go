package main

import (
	"fmt"
	"telehorn/config"

	"github.com/gin-gonic/gin"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/github"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	InitRouting(router)
	router.Run(appPort()) // listen and serve on 0.0.0.0:appPort
}

func appPort() string {
	configStruct := config.Get()
	return fmt.Sprintf("0.0.0.0:%s", configStruct.Port)
}
