package trips

type Trip struct {
	quantity int
}

func (t *Trip) GetQuantity() int {
	return t.quantity
}

func (t *Trip) SetQuantity(v int) {
	t.quantity = v
}

func New() *Trip {
	return &Trip{}
}
