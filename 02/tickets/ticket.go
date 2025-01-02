package tickets

type Ticket interface {
	GetQuantity() int
	SetQuantity(v int)
	// DiscountedForUser(userID int)
}
