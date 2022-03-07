package storage

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db   *gorm.DB
	once sync.Once
)

// singleton para open
func NewPostgresDB() {
	// lo que esta en la funcion anonima ,se va a ejecutar una sola vez , aunque ejecutemos NewPostgresDB varias veces
	once.Do(func() {
		var err error
		connStr := "postgres://postgres:nobunaga@localhost:5432/golang?sslmode=disable"
		db, err = gorm.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Cant open db: %v", err)
		}

		fmt.Println("Conectado a postgres")
	})
}

// Pool return a unique instance of db
func DB() *gorm.DB {
	return db
}
