package tripTicketRepo

import "github.com/nelsonsaake/learn-factory-design-pattern/02/tickets"

type TripTicketRepo struct{}

func (r *TripTicketRepo) Search(key string) ([]tickets.Ticket, error) {
	return search(key), nil
}

func New() *TripTicketRepo {
	return &TripTicketRepo{}
}
