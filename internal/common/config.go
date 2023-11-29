package common

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	Port     int64
	LogLevel int

	NamesClient struct {
		BaseUrl string
	}

	JokesClient struct {
		BaseUrl string
	}
}

func (c *config) load() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	logLevelStr := os.Getenv("LOG_LEVEL")
	logLevel, err := strconv.Atoi(logLevelStr)
	if err != nil {
		panic("invalid log level configured: " + logLevelStr)
	}
	c.LogLevel = logLevel

	portStr := os.Getenv("PORT")
	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		panic("invalid port configured: " + portStr)
	}
	c.Port = port

	c.NamesClient.BaseUrl = os.Getenv("NAMES_BASE_URL")
	c.JokesClient.BaseUrl = os.Getenv("JOKES_BASE_URL")
}

var internalConfig *config

func Config() *config {
	if internalConfig == nil {
		internalConfig = &config{}
		internalConfig.load()
	}

	return internalConfig
}
