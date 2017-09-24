package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

// ConnectionSettings - содержит конфигурационные данные подключения к БД
type ConnectionSettings struct {
	Adapter  string
	User     string
	Password string
	Database string
	Pool     int
	Sslmode  string
	Host     string
	Port     string
}

type DbVersion struct {
	versionRank   int64  `db:"version_rank"`
	installedRank int64  `db:"installed_rank"`
	Version       int64  `db:"version"`
	Description   string `db:"description"`
	migrationType string `db:"type"`
}

// OpenConnect - Возвращет переменную подключения к бд
// Для закрытия нужно использовать db_connect.Close()
func OpenConnect() (*sqlx.DB, error) {
	cs := settings()
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cs.Host, cs.Port, cs.User, cs.Password, cs.Database, cs.Sslmode)
	db, err := sqlx.Connect(cs.Adapter, connString)
	if err != nil {
		log.Fatalln(err)
	}
	return db, err
}

// Version - возвращает текующую версию бд
func Version() (int64, error) {
	db, err := OpenConnect()
	check(err)

	versionInfo := DbVersion{}

	err = db.Get(
		&versionInfo,
		"SELECT * FROM schema_migrations ORDER BY version DESC LIMIT(1)",
	)
	check(err)

	version := versionInfo.Version

	return version, err
}

// Заполняет и возвращает структуру настроек
func settings() ConnectionSettings {
	err := godotenv.Load()
	check(err)

	poolSize, err := strconv.Atoi(os.Getenv("DB_POOL_SIZE"))
	check(err)

	// Создаем структуру настроек подключения
	settings := ConnectionSettings{
		Adapter:  "postgres",
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
		Pool:     poolSize,
		Sslmode:  os.Getenv("DB_SSL_MODE"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT")}

	// Возвращаем структуру
	return settings
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}
