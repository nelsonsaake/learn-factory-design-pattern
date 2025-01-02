package main

import (
	"github.com/nelsonsaake/learn-factory-design-pattern/02/tickets"
)

type NewInventoryRequestItem struct {
	productType tickets.Type
	quantity    int
}

var (
	inventoryList = []NewInventoryRequestItem{
		{
			productType: tickets.Event,
			quantity:    21,
		},
		{
			productType: tickets.Trip,
			quantity:    53,
		},
	}
)

func InventoryManager() {

	cout("new inventory ...")

	cout("===")

	for _, v := range inventoryList {

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
