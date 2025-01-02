package eventTicketRepo

import (
	"strings"

	"github.com/nelsonsaake/learn-factory-design-pattern/02/tickets"
	"github.com/nelsonsaake/learn-factory-design-pattern/02/tickets/events"
)

var db = []events.Event{
	{Name: "Chale Wote Street Art Festival", Quantity: 1000},
	{Name: "Sankofa Festival", Quantity: 800},
	{Name: "Accra International Marathon", Quantity: 1500},
	{Name: "Dumsor Free Concert", Quantity: 500},
	{Name: "Ghana Music Awards", Quantity: 1200},
	{Name: "Bakatue Festival", Quantity: 700},
	{Name: "Homowo Festival", Quantity: 600},
	{Name: "Ghana Fashion Week", Quantity: 300},
	{Name: "Wli Waterfalls Festival", Quantity: 400},
	{Name: "Ghana Independence Day Parade", Quantity: 1500},
}

func search(key string) []tickets.Ticket {
	var (
		k   = strings.ToLower(key)
		res = []tickets.Ticket{}
	)
	for _, v := range db {
		scope := strings.ToLower(v.Name)
		if strings.Contains(scope, k) {
			res = append(res, &v)
		}
	}
	return res
}
