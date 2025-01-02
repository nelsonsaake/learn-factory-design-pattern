package main

import (
	"fmt"

	"github.com/nelsonsaake/learn-factory-design-pattern/02/ticketRepo"
	"github.com/nelsonsaake/learn-factory-design-pattern/02/tks"
)

type Query struct {
	TicketType tks.TicketType
	Key        string
}

var (
	queries = []Query{
		{
			TicketType: tks.EventTicket,
			Key:        "accra",
		},
		{
			TicketType: tks.TripTicket,
			Key:        "tarkwa",
		},
		{
			TicketType: tks.TripTicket,
			Key:        "accra",
		},
	}
)

func QueryManager() {

	cout("inventory query ...")

	cout("\n")

	for _, query := range queries {

		cout("key:", query.Key)

		repo, err := ticketRepo.New(query.TicketType)
		if err != nil {
			die(err)
		}

		tickets, err := repo.Search(query.Key)
		if err != nil {
			die(err)
		}

		cout("results:", len(tickets))
		cout("---")
		for i, ticket := range tickets {
			cout("["+fmt.Sprint(i+1)+"]", ticket)
		}

		cout("\n")
	}
}
