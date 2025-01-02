package tripTicketRepo

import (
	"strings"

	"github.com/nelsonsaake/learn-factory-design-pattern/02/tickets"
	"github.com/nelsonsaake/learn-factory-design-pattern/02/tickets/trips"
)

var db = []trips.Ticket{
	{From: "Accra", To: "Kumasi", Quantity: 50},
	{From: "Tamale", To: "Bolgatanga", Quantity: 30},
	{From: "Takoradi", To: "Cape Coast", Quantity: 20},
	{From: "Tema", To: "Ho", Quantity: 40},
	{From: "Wa", To: "Sunyani", Quantity: 25},
	{From: "Koforidua", To: "Nkawkaw", Quantity: 35},
	{From: "Sekondi", To: "Tarkwa", Quantity: 15},
	{From: "Elmina", To: "Winneba", Quantity: 10},
	{From: "Bawku", To: "Navrongo", Quantity: 12},
	{From: "Techiman", To: "Dormaa Ahenkro", Quantity: 18},
	{From: "Obuasi", To: "Bekwai", Quantity: 28},
	{From: "Agona Swedru", To: "Kasoa", Quantity: 22},
	{From: "Mampong", To: "Ejura", Quantity: 16},
	{From: "Nkawkaw", To: "Konongo", Quantity: 14},
	{From: "Sogakope", To: "Akatsi", Quantity: 19},
	{From: "Aflao", To: "Denu", Quantity: 13},
	{From: "Dome", To: "Madina", Quantity: 11},
	{From: "Nkwanta", To: "Dambai", Quantity: 7},
	{From: "Anloga", To: "Keta", Quantity: 9},
	{From: "Paga", To: "Navrongo", Quantity: 21},
}

func search(key string) []tickets.Ticket {
	var (
		k   = strings.ToLower(key)
		res = []tickets.Ticket{}
	)
	for _, v := range db {
		scope := strings.ToLower(v.From + v.To)
		if strings.Contains(scope, k) {
			res = append(res, &v)
		}
	}
	return res
}
