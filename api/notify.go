package api

import (
	"time"
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"telehorn/config"
)

import (
	db_supp "telehorn/db"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Struct JSON for creation notify
type NotifyCreationJSON struct {
	TypeID      int64     `json:"type_id" binding:"required"` // Type id from references(table notification_types)
	Title       string    `json:"title"   binding:"required"`
	Text        string    `json:"chat_id" binding:"required"`      // Chat for notification
	SecretToken string    `json:"secret_token" binding:"required"` // Access token for notify
	Tags        string    `json:"tags"`                            // Tags
	Time        time.Time `json:"notify_time"`                     // Time of notification
}

type Notification struct {
	ID          int64  `db:"id"`
	SubscribeID int64  `db:"subscribe_id"`
	TypeID      int64  `db:"type_id"`
	Title       string `db:"title"`
	Text        string `db:"text"`
	Tags        string `db:"tags"`
}

// Gin Handler for notification create
func CreateNotificationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var creationJson NotifyCreationJSON
		c.BindJSON(&creationJson)

		createdSubscribe, err := createNotification(c, creationJson)
		if err != nil {
			log.Panic(err)
		}

		pushNotification(createdSubscribe)

		c.JSON(201, gin.H{
			"OK": true,
		})
	}
}

func createNotification(c *gin.Context, creationJSON NotifyCreationJSON) (Notification, error) {
	db, err := db_supp.OpenConnect()
	defer db.Close()
	createdNotification := Notification{}

	if err != nil {
		log.Fatalln(err)
		return createdNotification, err
	}

	subscribe, err := GetSubscribeByToken(creationJSON.SecretToken)

	insertQuery := fmt.Sprintf(
		"INSERT INTO notifications (subscribe_id, type_id, title, text, tags) "+
			"VALUES (%b, %b, %s, %s, %s)",
		subscribe.ID, creationJSON.TypeID, creationJSON.Title,
		creationJSON.Text, creationJSON.Tags,
	)

	tx := db.MustBegin()
	db.MustExec(insertQuery)
	tx.Commit()

	selectQuery := fmt.Sprintf(
		"SELECT * FROM notifications WHERE subscribe_id = %b AND " +
			"type_id = %b AND title = %s ORDER BY id DESC",
				subscribe.ID, creationJSON.TypeID, creationJSON.Title,
	)
	db.Get(&createdNotification, selectQuery)
	return createdNotification, nil
}

func pushNotification(notification Notification) {
	bot, err := tgbotapi.NewBotAPI(config.BotKey())
	if err != nil {
		log.Panic(err)
	}

	subscribe, err := GetSubscribeByID(notification.SubscribeID)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	msg := tgbotapi.NewMessage(subscribe.ChatID, notification.Text)
	bot.Send(msg)
}
