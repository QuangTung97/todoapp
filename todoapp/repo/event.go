package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"todoapp/lib/dblib"
	"todoapp/todoapp/event/core"
	"todoapp/todoapp/model"
	"todoapp/todoapp/types"
)

// EventRepository ...
type EventRepository struct {
	db *sqlx.DB
}

var _ core.Repository = &EventRepository{}

// NewEventRepository ...
func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func modelEventsToCore(events []model.Event) []core.Event {
	result := make([]core.Event, 0, len(events))
	for _, e := range events {
		event := core.Event(types.EventFromModel(e))
		result = append(result, event)
	}
	return result
}

var getLastEventsQuery = dblib.NewQuery(`
SELECT e.id, e.sequence, e.data, e.created_at FROM (
	SELECT id, sequence, data, created_at FROM todo_events
	WHERE sequence IS NOT NULL
	ORDER BY sequence DESC
	LIMIT ?
) e ORDER BY sequence ASC
`)

// GetLastEvents ...
func (r *EventRepository) GetLastEvents(limit uint64) ([]core.Event, error) {
	var events []model.Event
	err := r.db.Select(&events, getLastEventsQuery, limit)
	if err != nil {
		return nil, err
	}
	return modelEventsToCore(events), nil
}

var getEventsFromSequenceQuery = dblib.NewQuery(`
SELECT id, sequence, data, created_at FROM todo_events
WHERE sequence IS NOT NULL AND sequence >= ?
ORDER BY sequence ASC
LIMIT ?
`)

// GetEventsFromSequence ...
func (r *EventRepository) GetEventsFromSequence(seq uint64, limit uint64) ([]core.Event, error) {
	var events []model.Event
	err := r.db.Select(&events, getEventsFromSequenceQuery, seq, limit)
	if err != nil {
		return nil, err
	}
	return modelEventsToCore(events), nil
}

var getUnprocessedEventsQuery = dblib.NewQuery(`
SELECT id, data, created_at FROM todo_events
WHERE sequence IS NULL
ORDER BY id ASC
LIMIT ?
`)

// GetUnprocessedEvents ...
func (r *EventRepository) GetUnprocessedEvents(limit uint64) ([]core.Event, error) {
	var events []model.Event
	err := r.db.Select(&events, getUnprocessedEventsQuery, limit)
	if err != nil {
		return nil, err
	}
	return modelEventsToCore(events), nil
}

var getLastSequenceQuery = dblib.NewQuery(`
SELECT sequence FROM todo_publishers
WHERE id = ?
`)

// GetLastSequence ...
func (r *EventRepository) GetLastSequence(id core.PublisherID) (uint64, error) {
	var result uint64
	err := r.db.Get(&result, getLastSequenceQuery, id)
	if err != nil {
		return 0, err
	}
	return result, nil
}

var saveLastSequence = dblib.NewQuery(`
UPDATE todo_publishers SET sequence = ? WHERE id = ?
`)

// SaveLastSequence ...
func (r *EventRepository) SaveLastSequence(id core.PublisherID, seq uint64) error {
	_, err := r.db.Exec(saveLastSequence, seq, id)
	return err
}

var updateSequencesQuery = `
INSERT INTO todo_events (id, sequence, data)
VALUES %s AS new
ON DUPLICATE KEY UPDATE sequence = new.sequence
`

var _ = dblib.NewQuery(fmt.Sprintf(updateSequencesQuery, "(?, ?, '')"))

// UpdateSequences ...
func (r *EventRepository) UpdateSequences(events []core.Event) error {
	if len(events) == 0 {
		return nil
	}

	var buf strings.Builder
	args := make([]interface{}, 0, 2*len(events))
	for i, e := range events {
		if i == 0 {
			buf.WriteString("(?, ?, '')")
		} else {
			buf.WriteString(",(?, ?, '')")
		}
		args = append(args, e.ID, e.Sequence)
	}

	query := fmt.Sprintf(updateSequencesQuery, buf.String())
	_, err := r.db.Exec(query, args...)
	return err
}
