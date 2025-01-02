package tickets

import (
	"github.com/nelsonsaake/learn-factory-design-pattern/02/tickets/events"
	"github.com/nelsonsaake/learn-factory-design-pattern/02/tickets/trips"
	"github.com/nelsonsaake/learn-factory-design-pattern/02/tks"
)

func New(t tks.TicketType) (Ticket, error) {
	switch t {
	case tks.EventTicket:
		return events.New(), nil
	case tks.TripTicket:
		return trips.New(), nil
	default:
		return nil, tks.UnsupportedType
	}
}
