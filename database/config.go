package database

import (
	"os"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	connection string
	host       string
	port       string
	name       string
	user       string
	password   string
}

func DatabaseConfig(configList *[]dbConfig) error {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// Initialize database configuration
	pgsql := dbConfig{
		connection: os.Getenv("DB_CONNECTION"),
		host:       os.Getenv("DB_HOST"),
		port:       os.Getenv("DB_PORT"),
		name:       os.Getenv("DB_DATABASE"),
		user:       os.Getenv("DB_USERNAME"),
		password:   os.Getenv("DB_PASSWORD"),
	}

	// set mysql database config
	mysql := dbConfig{
		connection: os.Getenv("DB_CONNECTION"),
		host:       os.Getenv("DB_HOST"),
		port:       os.Getenv("DB_PORT"),
		name:       os.Getenv("DB_DATABASE"),
		user:       os.Getenv("DB_USERNAME"),
		password:   os.Getenv("DB_PASSWORD"),
	}

	// Slice untuk menyimpan konfigurasi database
	*configList = []dbConfig{pgsql, mysql}

	return nil
}
