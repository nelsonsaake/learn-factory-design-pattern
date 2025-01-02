package tickets

import (
	"errors"

	"github.com/nelsonsaake/learn-factory-design-pattern/02/tickets/events"
	"github.com/nelsonsaake/learn-factory-design-pattern/02/tickets/trips"
)

func New(t Type) (Ticket, error) {
	switch t {
	case Event:
		return events.New(), nil
	case Trip:
		return trips.New(), nil
	default:
		return nil, errors.New("unsupport product type provided")
	}
}
