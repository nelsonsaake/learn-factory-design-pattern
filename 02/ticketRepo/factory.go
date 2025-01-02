package ticketRepo

import (
	"github.com/nelsonsaake/learn-factory-design-pattern/02/ticketRepo/eventTicketRepo"
	"github.com/nelsonsaake/learn-factory-design-pattern/02/ticketRepo/tripTicketRepo"
	"github.com/nelsonsaake/learn-factory-design-pattern/02/tks"
)

func New(t tks.TicketType) (TicketRepo, error) {
	switch t {
	case tks.EventTicket:
		return eventTicketRepo.New(), nil
	case tks.TripTicket:
		return tripTicketRepo.New(), nil
	default:
		return nil, tks.UnsupportedType
	}
}
