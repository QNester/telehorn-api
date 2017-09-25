package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct - Структура конфигурации
type Config struct {
	// HTTP configuration
	Port string
	// Environment
	Environment string
}

// configInstance - Инстанс структуры конфигурации
var configInstance Config

// Get - Получение объекта конфигурации
func Get() Config {
	emptyConfig := Config{}
	if configInstance == emptyConfig {
		setUp()
	}
	return configInstance
}

func setUp() {
	confStruct := Config{
		port(),
		environment(),
	}
	configInstance = confStruct
}

func envKey(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

// IsProduction - Return true if environment is production
func (c Config) IsProduction() bool {
	return c.Environment == "production"
}

// return current environment
func environment() string {
	return envKey("APP_ENVIRONMENT")
}

// return current port
func port() string {
	return envKey("APP_PORT")
}

func BotKey() string {
	return envKey("BOT_KEY")
}
