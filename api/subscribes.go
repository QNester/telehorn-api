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
		fmt.Printf("Request from user with id %b to add chat %b", creationJson.UserID, creationJson.ChatID)

		createdSubscribe := createRow(c, creationJson)

		c.JSON(201, gin.H{
			"user_id": createdSubscribe.UserID,
			"chat_id": createdSubscribe.ChatID,
			"secret_token": createdSubscribe.SecretToken,
			"created_at": createdSubscribe.CreatedAt,
		})
	}
}

func randomToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func createRow(c *gin.Context, creationJSON CreateJSON) Subscriber{
	db, err := db_supp.OpenConnect()
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON(500, gin.H{
			"db_status": "Fail",
		})
	}

	insertQuery := fmt.Sprintf("INSERT INTO subscribers (user_id, chat_id, secret_token, created_at) " +
			"VALUES (%b, %b, '%s', $1)", creationJSON.UserID, creationJSON.ChatID, randomToken())

	fmt.Println(insertQuery)
	tx := db.MustBegin()
	db.MustExec(insertQuery, time.Now())
	tx.Commit()

	selectQuery := fmt.Sprintf(
		"SELECT * FROM subscribers WHERE user_id = %b AND chat_id = %b",
			creationJSON.UserID, creationJSON.ChatID,
	)
	fmt.Println(selectQuery)
	createdSubscribe := Subscriber{}
	db.Get(&createdSubscribe, selectQuery)
	fmt.Printf("%#v\n", creationJSON)

	return createdSubscribe
}
