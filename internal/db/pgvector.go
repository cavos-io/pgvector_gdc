package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// DB is the global database connection
var DB *sql.DB

// InitDB initializes the database connection and prepares for vector queries
func InitDB() error {
	// Get database connection string from environment variables
	connStr := os.Getenv("PG_DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:postgrespassword@postgres:5432/postgres?sslmode=disable"
	}

	// Connect to the database
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// Verify connection
	if err = DB.Ping(); err != nil {
		return err
	}

	// Log success
	log.Println("Connected to the database successfully")
	return nil
}
