package api

import (
	db_supp "telehorn/db"
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
	"crypto/rand"
	"time"
)

// Struct JSON for creation subscribe
type CreateJSON struct {
	UserID int64 `json:"user_id" binding:"required"` // User who registrate chat for bot
	ChatID int64 `json:"chat_id" binding:"required"` // Chat for notification
}

// Struct Subscriber
type Subscriber struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	ChatID      int64     `db:"chat_id"`
	SecretToken string    `db:"secret_token"`
	CreatedAt   time.Time `db:"created_at"`
}

func CreateSubscribeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var creationJson CreateJSON
		c.BindJSON(&creationJson)

		log.Printf("Request from user with id %b to add chat %b", creationJson.UserID, creationJson.ChatID)

		db, err := db_supp.OpenConnect()
		defer db.Close()

		//newSubscriber := createRow()
		fmt.Println(randomToken())

		log.Fatalln(err)

		if err != nil {
			c.JSON(500, gin.H{
				"db_status": "Fail",
			})
		}
	}

	return func(c *gin.Context) {
		c.JSON(500, gin.H{
			"secret_token": "OK",
		})
	}
}

func randomToken() string {
	b := make([]byte, 64)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func createRow() {

}
