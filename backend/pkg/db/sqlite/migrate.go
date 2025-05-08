package sqlite

import (
    "os"
    "log"
    "path/filepath"
    "strings"
)

func ApplyMigrations() {
    migrationDir := "./backend/pkg/db/migrations/sqlite"

    // Walk through the migration directory
    err := filepath.Walk(migrationDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Apply only .up.sql files
        if !info.IsDir() && strings.HasSuffix(info.Name(), ".up.sql") {
            sqlBytes, err := os.ReadFile(path)
            if err != nil {
                return err
            }

            _, err = DB.Exec(string(sqlBytes))
            if err != nil {
                log.Fatalf("Failed to apply migration: %s, error: %v", path, err)
            }

            log.Printf("Applied migration: %s", path)
        }

        return nil
    })

    if err != nil {
        log.Fatal("Failed to apply migrations:", err)
    }
}
