package db

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewSQL creates and SQL connection using environment variables
// to configure.
func NewSQL() (*sqlx.DB, error) {
	host := strings.TrimSpace(os.Getenv("POSTGRES_HOST"))
	port := strings.TrimSpace(os.Getenv("POSTGRES_PORT"))
	user := strings.TrimSpace(os.Getenv("POSTGRES_USER"))
	password := strings.TrimSpace(os.Getenv("POSTGRES_PASSWORD"))
	db := strings.TrimSpace(os.Getenv("POSTGRES_DB"))

	info := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		db,
	)

	return sqlx.Connect("postgres", info)
}

func CreateTables(db *sqlx.DB) error {
	schemaInitString := `
		CREATE TABLE IF NOT EXISTS event (
				id UUID PRIMARY KEY NOT NULL,
				name TEXT NOT NULL,
				description TEXT NOT NULL,
				start_date TIMESTAMP WITH TIME ZONE NOT NULL,
				end_date TIMESTAMP WITH TIME ZONE NOT NULL,
				location TEXT NOT NULL,
				is_paid BOOLEAN NOT NULL,
				total_cost NUMERIC,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
				last_update TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
		);

		CREATE TABLE IF NOT EXISTS attendee (
				id UUID PRIMARY KEY NOT NULL,
				event_id UUID NOT NULL,
				name TEXT NOT NULL,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
				last_update TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
				FOREIGN KEY (event_id) REFERENCES event(id)
		);
	`

	_, err := db.Exec(schemaInitString)

	return err
}

func SeedDataBase(db *sqlx.DB) error {
	fileAsBytes, err := os.ReadFile("seed.sql")

	if err != nil {
		slog.Error("Could not feed seed.sql in root of project")
		return err
	}

	_, err = db.Exec(string(fileAsBytes))

	if err != nil {
		slog.Error("Could not run seed script")
		return err
	}

	return nil
}
