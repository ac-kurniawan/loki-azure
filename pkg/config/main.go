package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PSQL_HOST     string
	PSQL_PORT     string
	PSQL_USERNAME string
	PSQL_PASSWORD string
	PSQL_DB_NAME  string

	ServerPort string
}

func (c *Config) DefaultConfig() {
	c.PSQL_HOST = "localhost"
	c.PSQL_PORT = "5432"
	c.PSQL_USERNAME = "root"
	c.PSQL_PASSWORD = "root"
	c.PSQL_DB_NAME = "postgres"

	c.ServerPort = "3222"
}

func (c *Config) GetConfig() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Printf("Error loading .env file")
		c.DefaultConfig()
	}

	// PSQL config
	c.PSQL_HOST = os.Getenv("PSQL_HOST")
	c.PSQL_PORT = os.Getenv("PSQL_PORT")
	c.PSQL_USERNAME = os.Getenv("PSQL_USERNAME")
	c.PSQL_PASSWORD = os.Getenv("PSQL_PASSWORD")
	c.PSQL_DB_NAME = os.Getenv("PSQL_DB_NAME")

	c.ServerPort = "3222"
}
