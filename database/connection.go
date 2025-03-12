package database

import (
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	err  error
	once sync.Once
)

var dbConnections = make(map[string]*gorm.DB)

func Init(scheme string) (*gorm.DB, error) {
	configList := []dbConfig{}

	err := DatabaseConfig(&configList)
	if err != nil {
		log.Fatal("Error loading database configuration:", err)
	}

	// Looping untuk mencari database yang sesuai dengan scheme
	for _, config := range configList {

		if config.connection == scheme {
			// Cek apakah koneksi sudah ada
			if db, exists := dbConnections[scheme]; exists {
				return db, nil
			}

			// Membuat koneksi baru
			db, err := ConnectDB(config.connection, config.host, config.port, config.name, config.user, config.password, 3)
			if err != nil {
				return nil, err
			}

			// Simpan koneksi ke map
			dbConnections[scheme] = db
			return db, nil
		}
	}

	return nil, fmt.Errorf("no matching database found for scheme: %s", scheme)
}

// ConnectDB menghubungkan ke database berdasarkan tipe yang dipilih
func ConnectDB(dbType string, dbHost string, dbPort string, dbName string, dbUser string, dbPass string, attempt int) (*gorm.DB, error) {
	once.Do(func() {
		var dsn string
		var dialector gorm.Dialector // <-- Ini untuk memilih driver sesuai dbType

		// Buat DSN berdasarkan tipe database
		switch dbType {
		case "pgsql":
			dsn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				dbHost, dbPort, dbUser, dbPass, dbName,
			)
			dialector = postgres.Open(dsn)

		case "mysql":
			dsn = fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				dbUser, dbPass, dbHost, dbPort, dbName,
			)
			dialector = mysql.Open(dsn) // <-- Gunakan driver MySQL

		case "sqlite":
			dsn = fmt.Sprintf("%s.db", dbName)
			dialector = sqlite.Open(dsn)

		default:
			err = fmt.Errorf("❌ Database type '%s' not supported", dbType)
			return
		}

		// Coba koneksi ulang jika gagal
		maxRetries := attempt
		for i := 1; i <= maxRetries; i++ {
			db, err = gorm.Open(dialector, &gorm.Config{})
			if err == nil {
				fmt.Println("✅ Database Connected Successfully!")
				return
			}

			log.Printf("⚠️ Attempt %d: Failed to connect to database: %v\n", i, err)
			time.Sleep(2 * time.Second) // Tunggu sebelum mencoba lagi
		}

		log.Fatalf("❌ Failed to connect to database after %d attempts: %v", maxRetries, err)
	})

	return db, err
}
