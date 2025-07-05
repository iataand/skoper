package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/iataand/skoper/internal/utils"
	_ "github.com/lib/pq"
)

func createTables(db *sql.DB) error {
	_, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`)
	if err != nil {
		return fmt.Errorf("failed to create extension: %w", err)
	}

	programsTable := `
	CREATE TABLE IF NOT EXISTS programs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		handle TEXT NOT NULL UNIQUE,
		handleApiUrl TEXT NOT NULL 
	);`

	scopesTable := `
	CREATE TABLE IF NOT EXISTS scopes (
		id TEXT PRIMARY KEY,
		program_id UUID NOT NULL REFERENCES programs(id) ON DELETE CASCADE,
		asset_type TEXT NOT NULL,
		asset_identifier TEXT NOT NULL,
		eligible_for_bounty BOOLEAN NOT NULL,
		eligible_for_submission BOOLEAN NOT NULL,
		instruction TEXT,
		max_severity TEXT,
		created_at TIMESTAMPTZ,
		updated_at TIMESTAMPTZ,
		confidentiality_requirement TEXT,
		integrity_requirement TEXT,
		availability_requirement TEXT
	);`

	if _, err := db.Exec(programsTable); err != nil {
		return fmt.Errorf("failed to create 'programs' table: %w", err)
	}

	if _, err := db.Exec(scopesTable); err != nil {
		return fmt.Errorf("failed to create 'scopes' table: %w", err)
	}

	log.Println("🗃️ Tables 'programs' and 'scopes' created or already exist")
	return nil
}

func InitDbConnection() (*sql.DB, error) {
	host, port, user, password, dbname := utils.LoadDbEnvVariables()

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error opening DB connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging DB: %w", err)
	}

	log.Println("✅ Successfully connected to the database")

	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("error creating tables: %w", err)
	}

	return db, nil
}
