package ticketRepo

import "github.com/nelsonsaake/learn-factory-design-pattern/02/tickets"

type TicketRepo interface {
	Search(key string) ([]tickets.Ticket, error)
}
