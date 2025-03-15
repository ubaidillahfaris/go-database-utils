package database

import (
	"github.com/joho/godotenv"
)

// dbConfig adalah struktur untuk konfigurasi database
type dbConfig struct {
	connection string
	host       string
	port       string
	name       string
	user       string
	password   string
}

// DatabaseConfig mengisi konfigurasi database dari environment variables
func DatabaseConfig(configList *[]dbConfig) error {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		return err
	}

	// PostgreSQL
	pgsql := dbConfig{
		connection: "pgsql",
		host:       "localhost",
		port:       "5432",
		name:       "parfume_pos_app",
		user:       "faris",
		password:   "12345faris",
	}

	// Konfigurasi MySQL
	//mysql := dbConfig{
	//	connection: "mysql",
	//	host:       "localhost",
	//	port:       "3306",
	//	name:       "perfume_old",
	//	user:       "root",
	//	password:   "root",
	//}

	*configList = []dbConfig{pgsql, /**mysql**/}

	return nil
}
