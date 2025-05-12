package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

func HelpeMegration() {
	dbPath := "./backend/social-network.db"
	migrationsDir := "./backend/pkg/db/migrations/sqlite"

	// Connect to the SQLite database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Apply the migrations
	err = applyMigrations(db, migrationsDir)
	if err != nil {
		log.Fatal("Failed to apply migrations:", err)
	}

	log.Println("Migrations applied successfully.")
}

func applyMigrations(db *sql.DB, dir string) error {
	// err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
	//     if err != nil {
	//         return err
	//     }

	//     // Only apply .up.sql files
	//     if strings.HasSuffix(info.Name(), ".up.sql") {
	//         log.Println("Applying migration:", info.Name())
	//         content, err := os.ReadFile(path)
	//         if err != nil {
	//             return err
	//         }

	//         _, err = db.Exec(string(content))
	//         if err != nil {
	//             return err
	//         }
	//     }

	//     return nil
	// })
	migrations := &migrate.FileMigrationSource{
		Dir: dir,
	}
	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	fmt.Println("number: ", n)
	return err
}
