package infrastructure

import (
	"os"
	"strconv"
)

type Config struct {
	AMQ_USERNAME string
	AMQ_PASSWORD string
	AMQ_HOST     string
	AMQ_VHOST    string
	AMQ_PORT     int
}

func NewConfigFromEnvironmental() (*Config, error) {
	port := os.Getenv("AMQ_PORT")
	portAsInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	return &Config{
		AMQ_USERNAME: os.Getenv("AMQ_USERNAME"),
		AMQ_PASSWORD: os.Getenv("AMQ_PASSWORD"),
		AMQ_HOST:     os.Getenv("AMQ_HOST"),
		AMQ_VHOST:    os.Getenv("AMQ_VHOST"),
		AMQ_PORT:     portAsInt,
	}, nil
}
