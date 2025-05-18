package sqlite

import (
    "database/sql"
    "log"
    "os"
    "path/filepath"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
    dbPath := "./backend/social-network.db"

    // Ensure the directory exists
    dir := filepath.Dir(dbPath)
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        log.Println("Database directory not found, creating it...")
        if err := os.MkdirAll(dir, 0755); err != nil {
            log.Fatal("Failed to create database directory:", err)
        }
    }

    // Create the database file if it doesn't exist
    if _, err := os.Stat(dbPath); os.IsNotExist(err) {
        log.Println("Database file not found, creating a new one...")
        file, err := os.Create(dbPath)
        if err != nil {
            log.Fatal("Failed to create database file:", err)
        }
        file.Close()
    }

    var err error
    DB, err = sql.Open("sqlite3", dbPath)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Test the connection
    if err := DB.Ping(); err != nil {
        log.Fatal("Failed to ping database:", err)
    }

    log.Println("Database connection established.")
}
