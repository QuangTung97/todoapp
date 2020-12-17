package model

import (
	"database/sql"
	"time"
)

// EventID ...
type EventID uint64

// Event ...
type Event struct {
	ID        EventID       `db:"id"`
	Sequence  sql.NullInt64 `db:"sequence"`
	Data      string        `db:"data"`
	CreatedAt time.Time     `db:"created_at"`
}
