package database

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/joho/godotenv"
)

// GenerateConfig membuat file database.go berdasarkan template
func GenerateConfig() error {
	// Tentukan lokasi template dan output
	templatePath := filepath.Join("database", "tmpl", "config.tmpl")
	outputPath := filepath.Join("config", "database.go")

	// Baca file template
	tmplContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("gagal membaca template: %w", err)
	}

	// Parse template
	tmpl, err := template.New("config").Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("gagal memproses template: %w", err)
	}

	if err := godotenv.Load(); err != nil {
		return err
	}

	// Data default dengan dukungan environment variables
	data := map[string]string{
		"PG_CONNECTION":    getEnv("PG_CONNECTION", "pgsql"),
		"PG_HOST":          getEnv("PG_HOST", "127.0.0.1"),
		"PG_PORT":          getEnv("PG_PORT", "5432"),
		"PG_NAME":          getEnv("PG_DATABASE", "mydatabase"),
		"PG_USER":          getEnv("PG_USERNAME", "admin"),
		"PG_PASSWORD":      getEnv("PG_PASSWORD", "secret"),
		"MYSQL_CONNECTION": getEnv("MYSQL_CONNECTION", "mysql"),
		"MYSQL_HOST":       getEnv("MYSQL_HOST", "127.0.0.1"),
		"MYSQL_PORT":       getEnv("MYSQL_PORT", "3306"),
		"MYSQL_NAME":       getEnv("MYSQL_DATABASE", "mydatabase"),
		"MYSQL_USER":       getEnv("MYSQL_USERNAME", "root"),
		"MYSQL_PASSWORD":   getEnv("MYSQL_PASSWORD", "password"),
	}

	// Proses template dengan data
	var output bytes.Buffer
	if err := tmpl.Execute(&output, data); err != nil {
		return fmt.Errorf("gagal mengeksekusi template: %w", err)
	}

	// Buat folder jika belum ada
	if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
		return fmt.Errorf("gagal membuat direktori: %w", err)
	}

	// Simpan file konfigurasi
	if err := os.WriteFile(outputPath, output.Bytes(), 0644); err != nil {
		return fmt.Errorf("gagal menyimpan file konfigurasi: %w", err)
	}

	fmt.Println("Konfigurasi database berhasil dibuat di:", outputPath)
	return nil
}

// getEnv membaca variabel lingkungan atau mengembalikan nilai default jika tidak ada
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
