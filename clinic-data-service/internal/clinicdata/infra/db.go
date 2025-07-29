package infra

import (
	"database/sql"
	"fmt"
	"log"

	"clinic-data-service/config"
	_ "github.com/lib/pq"
)

// InitDB connects to PostgreSQL and returns the sql.DB instance
func InitDB(env *config.EnvConfig) *sql.DB {
	// Format DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DBHost,
		env.DBPort,
		env.DBUser,
		env.DBPass,
		env.DBName,
	)

	// Buka koneksi
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	// Cek koneksi
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	log.Println("Connected to PostgreSQL")
	return db
}
