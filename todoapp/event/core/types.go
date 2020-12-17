package core

import "todoapp/todoapp/types"

// Event ...
type Event types.Event

// SetSequenceImpl ...
func SetSequenceImpl(e Event, seq uint64) Event {
	e.Sequence = seq
	return e
}

// GetSequenceImpl ...
func GetSequenceImpl(e Event) uint64 {
	return e.Sequence
}
