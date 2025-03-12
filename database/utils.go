package database

import (
	"log"

	"gorm.io/gorm"
)

type QueryBuilder struct {
	db *gorm.DB
}

func DB(name string, scheme string) *QueryBuilder {

	if len(scheme) == 0 {
		scheme = "pgsql"
	}
	db, err := Init(scheme)
	if err != nil {
		log.Fatal(err)
	}
	return &QueryBuilder{db: db.Table(name)}
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
