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

// Struct Subscribe
type Subscribe struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	ChatID      int64     `db:"chat_id"`
	SecretToken string    `db:"secret_token"`
	CreatedAt   time.Time `db:"created_at"`
}

// Gin Handler for subscribe create
func CreateSubscribeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var creationJson CreateJSON
		c.BindJSON(&creationJson)

		createdSubscribe := createSubscribe(c, creationJson)

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

func createSubscribe(c *gin.Context, creationJSON CreateJSON) Subscribe{
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

	tx := db.MustBegin()
	db.MustExec(insertQuery, time.Now())
	tx.Commit()

	selectQuery := fmt.Sprintf(
		"SELECT * FROM subscribers WHERE user_id = %b AND chat_id = %b",
			creationJSON.UserID, creationJSON.ChatID,
	)
	createdSubscribe := Subscribe{}
	db.Get(&createdSubscribe, selectQuery)
	return createdSubscribe
}

func GetSubscribeByToken(secret_token string) (Subscribe, error) {
	foundSubscribe := Subscribe{}
	db, err := db_supp.OpenConnect()
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
		return foundSubscribe, err
	}

	selectQuery := fmt.Sprintf(
		"SELECT * FROM subscribers WHERE secret_token = %s",
			secret_token,
	)

	db.Get(&foundSubscribe, selectQuery)
	return foundSubscribe, nil
}

func GetSubscribeByID(id int64) (Subscribe, error) {
	foundSubscribe := Subscribe{}
	db, err := db_supp.OpenConnect()
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
		return foundSubscribe, err
	}

	selectQuery := fmt.Sprintf(
		"SELECT * FROM subscribers WHERE id = %s",
			id,
	)

	db.Get(&foundSubscribe, selectQuery)
	return foundSubscribe, nil
}
