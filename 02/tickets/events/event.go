package events

type Event struct {
	quantity int
}

func (e *Event) GetQuantity() int {
	return e.quantity
}

func (e *Event) SetQuantity(v int) {
	e.quantity = v
}

func New() *Event {
	return &Event{}
}
