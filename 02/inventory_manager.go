package main

import (
	"github.com/nelsonsaake/learn-factory-design-pattern/02/tickets"
	"github.com/nelsonsaake/learn-factory-design-pattern/02/tks"
)

type InventoryItem struct {
	productType tks.TicketType
	quantity    int
}

var (
	inventory = []InventoryItem{
		{
			productType: tks.EventTicket,
			quantity:    21,
		},
		{
			productType: tks.TripTicket,
			quantity:    53,
		},
	}
)

func InventoryManager() {

	cout("new inventory ...")

	cout("===")

	for _, v := range inventory {

		cout("product:", v.productType)

		p, err := tickets.New(v.productType)
		if err != nil {
			die(err)
		}

		cout("current quantity:", p.GetQuantity())

		cout("new inventory quantity:", v.quantity)

		p.SetQuantity(p.GetQuantity() + v.quantity)

		cout("updated quantity:", p.GetQuantity())

		cout("---")
	}
}
