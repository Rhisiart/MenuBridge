package config

import (
	"os"
	"strconv"
)

type Configuration struct {
	Port        uint
	DatabaseUrl string
}

func NewConfiguration() (*Configuration, error) {
	port, err := GetPort()

	if err != nil {
		return nil, err
	}

	return &Configuration{
		Port:        port,
		DatabaseUrl: os.Getenv("DATABASEURL"),
	}, nil
}

func GetPort() (uint, error) {
	portStr := os.Getenv("PORT")
	portEnv, err := strconv.Atoi(portStr)

	if err != nil {
		return 0, err
	}

	return uint(portEnv), nil
}
