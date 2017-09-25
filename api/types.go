package api

import (
	db_supp "telehorn/db"
	"log"
)

type NotificationType struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Img         string `db:"img_url"`
	Moderated   string `db:"moderated"`
}

func GetTypes() ([]NotificationType, error) {
	notificationTypes := []NotificationType{}

	db, err := db_supp.OpenConnect()
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
		return notificationTypes, err
	}

	db.Select(&notificationTypes, "SELECT * FROM notification_types WHERE moderated = true")

	return notificationTypes, nil
}
