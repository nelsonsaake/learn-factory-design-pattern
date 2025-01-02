package events

import "fmt"

type Event struct {
	Name     string
	Quantity int
}

func (e *Event) GetQuantity() int {
	return e.Quantity
}

func (e *Event) SetQuantity(v int) {
	e.Quantity = v
}

func (e Event) String() string {
	return fmt.Sprintf(e.Name)
}

func New() *Event {
	return &Event{}
}
