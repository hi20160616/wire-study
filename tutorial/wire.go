//go:build wireinject

package tutorial

import "github.com/google/wire"

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
