package main

import (
	db_supp "telehorn/db"
	"log"

	"github.com/gin-gonic/gin"
	"telehorn/api"
)

// InitRouting - load route handlers
func InitRouting(router *gin.Engine) {
	router.GET("/ping", hello())
	router.GET("/check_db", checkDb())
	router.POST("/subscribes", api.CreateSubscribeHandler())
	router.POST("/notify", api.CreateNotificationHandler())
}

func hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	}
}

func checkDb() gin.HandlerFunc {
	db, err := db_supp.OpenConnect()
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
		return func(c *gin.Context) {
			c.JSON(200, gin.H{
				"db_status": "Fail",
			})
		}
	}

	err = db.Ping()

	if err != nil {
		log.Fatalln(err)
		return func(c *gin.Context) {
			c.JSON(200, gin.H{
				"db_status": "Fail",
			})
		}
	}

	dbVersion := 1

	return func(c *gin.Context) {
		c.JSON(500, gin.H{
			"db_status":       "OK",
			"current_version": dbVersion,
		})
	}
}
