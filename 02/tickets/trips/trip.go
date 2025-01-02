package trips

import "fmt"

type Ticket struct {
	From     string
	To       string
	Quantity int
}

func (t *Ticket) GetQuantity() int {
	return t.Quantity
}

func (t *Ticket) SetQuantity(v int) {
	t.Quantity = v
}

func (t Ticket) String() string {
	return fmt.Sprintf("%v tickets, %v to %v", t.Quantity, t.From, t.To)
}

func New() *Ticket {
	return &Ticket{}
}
