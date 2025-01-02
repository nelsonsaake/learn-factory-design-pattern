package eventTicketRepo

import "github.com/nelsonsaake/learn-factory-design-pattern/02/tickets"

type EventTicketRepo struct{}

func (r *EventTicketRepo) Search(key string) ([]tickets.Ticket, error) {
	return search(key), nil
}

func New() *EventTicketRepo {
	return &EventTicketRepo{}
}
