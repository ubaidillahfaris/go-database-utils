package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type QueryBuilder struct {
	db *gorm.DB
}

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

// All mengambil semua data
func (qb *QueryBuilder) All(dest interface{}) error {
	return qb.db.Find(dest).Error
}

// First mengambil satu data pertama
func (qb *QueryBuilder) First(dest interface{}) error {
	return qb.db.First(dest).Error
}

// Where untuk filter query
func (qb *QueryBuilder) Where(query string, args ...interface{}) *QueryBuilder {
	qb.db = qb.db.Where(query, args...)
	return qb
}

// OrWhere menambahkan kondisi OR ke query
func (qb *QueryBuilder) OrWhere(query string, args ...interface{}) *QueryBuilder {
	qb.db = qb.db.Or(query, args...)
	return qb
}

// Limit membatasi jumlah hasil
func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.db = qb.db.Limit(limit)
	return qb
}
