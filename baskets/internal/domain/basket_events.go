package domain

const (
	BasketStartedEvent     = "baskets.BasketStarted"
	BasketItemAddedEvent   = "baskets.BasketItemAdded"
	BasketItemRemovedEvent = "baskets.BasketItemRemoved"
	BasketCanceledEvent    = "baskets.BasketCanceled"
	BasketCheckedOutEvent  = "baskets.BasketCheckedOut"
)

type BasketStarted struct {
	CustomerID string
}

type BasketItemAdded struct {
	Item Item
}

type BasketItemRemoved struct {
	ProductID string
	Quantity  int
}

type BasketCanceled struct{}

type BasketCheckedOut struct {
	PaymentID  string
	CustomerID string
	Items      map[string]Item
}
