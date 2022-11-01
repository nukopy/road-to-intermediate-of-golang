package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	ENV string
	PORT int
}

func New() *Config {
	return &Config{
		ENV: getEnv(),
		PORT: getPort(),
	}
}

func getEnv() string {
	ENV := os.Getenv("ENV")
	if ENV == "" {
		panic("ENV is not set.")
	}

	return ENV
}

func getPort() int {
	_PORT := os.Getenv("PORT")
	if _PORT == "" {
		panic("PORT is not set.")
	}
	PORT, err := strconv.Atoi(_PORT)
	if err != nil {
		msg := fmt.Sprintf("PORT cannot be converted to int: %s,\nERROR: %+v", _PORT, err)
		panic(msg)
	}

	return PORT
}