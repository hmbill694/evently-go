package model

import (
	"time"

	"github.com/google/uuid"
)

// Event matches the EVENT table
type Event struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	StartDate   time.Time `db:"start_date"`
	EndDate     time.Time `db:"end_date"`
	Location    string    `db:"location"`
	IsPaid      bool      `db:"is_paid"`
	TotalCost   float64   `db:"total_cost"` // Adjust the type as necessary based on the precision you need
	CreatedAt   time.Time `db:"created_at"`
	LastUpdate  time.Time `db:"last_update"`
}

// Attendee matches the ATTENDEE table
type Attendee struct {
	ID         uuid.UUID `db:"id"`
	EventID    uuid.UUID `db:"event_id"`
	Name       string    `db:"name"`
	CreatedAt  time.Time `db:"created_at"`
	LastUpdate time.Time `db:"last_update"`
}
