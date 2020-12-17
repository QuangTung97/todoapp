package types

import (
	"database/sql"
	"github.com/golang/protobuf/proto"
	"time"
	todoapp_rpc "todoapp-rpc/rpc/todoapp/v1"
	"todoapp/todoapp/model"
)

// Event ...
type Event struct {
	ID        model.EventID
	Sequence  uint64
	Data      *todoapp_rpc.Event
	CreatedAt time.Time
}

// ToModel ...
func (e Event) ToModel() model.Event {
	data, err := proto.Marshal(e.Data)
	if err != nil {
		panic(err)
	}

	return model.Event{
		ID: e.ID,
		Sequence: sql.NullInt64{
			Valid: true,
			Int64: int64(e.Sequence),
		},
		Data:      string(data),
		CreatedAt: e.CreatedAt,
	}
}

// EventFromModel ...
func EventFromModel(e model.Event) Event {
	var data *todoapp_rpc.Event
	err := proto.Unmarshal([]byte(e.Data), data)
	if err != nil {
		panic(err)
	}

	return Event{
		ID:        e.ID,
		Sequence:  uint64(e.Sequence.Int64),
		Data:      data,
		CreatedAt: e.CreatedAt,
	}
}
