package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// Struktur utama QueryBuilder
type QueryBuilder struct {
	db *gorm.DB
}

// Fungsi untuk inisialisasi database
func DB(name string, scheme ...string) (*QueryBuilder, error) {
	if len(scheme) > 1 {
		return nil, fmt.Errorf("Error: Too many parameters, just 2 parameters allowed")
	}

	dbScheme := "pgsql" // Default scheme
	if len(scheme) > 0 {
		dbScheme = scheme[0] // Ambil nilai pertama jika diberikan
	}
	db, err := Init(dbScheme)

	if err != nil {
		log.Fatal(err)
	}
	return &QueryBuilder{db: db.Table(name)}, nil
}

// Method untuk mengambil semua data
func (qb *QueryBuilder) All(dest interface{}) error {
	return qb.db.Find(dest).Error
}

// Method untuk mengambil satu data pertama
func (qb *QueryBuilder) First(dest interface{}) error {
	return qb.db.First(dest).Error
}

// Method untuk menambahkan filter WHERE
func (qb *QueryBuilder) Where(query string, args ...interface{}) *QueryBuilder {
	qb.db = qb.db.Where(query, args...)
	return qb
}

// Method untuk menambahkan filter OR WHERE
func (qb *QueryBuilder) OrWhere(query string, args ...interface{}) *QueryBuilder {
	qb.db = qb.db.Or(query, args...)
	return qb
}

// Method untuk menambahkan limit
func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.db = qb.db.Limit(limit)
	return qb
}
