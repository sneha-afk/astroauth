package store

import (
	"fmt"
	"io"
	"os"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() error {
	// Set up DB and create if does not exist
	var err error
	DB, err = sqlx.Open("sqlite3", "./store/astroauth.db")
	if err != nil {
		return fmt.Errorf("InitDB(): could not open DB: %v", err)
	}
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("InitDB(): could not ping DB: %v", err)
	}
	return nil
}

func CloseDB() error {
	if err := DB.Close(); err != nil {
		return fmt.Errorf("CloseDB: could not close DB: %v", err)
	}
	return nil
}

// Execute an .sql file given path
func ExecuteSQLFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("ExecuteFile: failed opening file: %w", err)
	}
	defer file.Close()

	schema, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("ExecuteFile: failed reading file: %w", err)
	}

	_, err = DB.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("ExecuteFile: failed to execute file: %w", err)
	}

	return nil
}
