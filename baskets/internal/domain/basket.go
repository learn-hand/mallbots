package domain

import (
	"github.com/learn-hand/mallbots/internal/ddd"
	"github.com/stackus/errors"
)

const BasketAggregate = "baskets.Basket"

var (
	ErrBasketHasNoItems         = errors.Wrap(errors.ErrBadRequest, "the basket has no items")
	ErrBasketCannotBeModified   = errors.Wrap(errors.ErrBadRequest, "the basket cannot be modified")
	ErrBasketCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the basket cannot be cancelled")
	ErrQuantityCannotBeNegative = errors.Wrap(errors.ErrBadRequest, "the item quantity cannot be negative")
	ErrBasketIDCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the basket id cannot be blank")
	ErrPaymentIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
	ErrCustomerIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
)

type Item struct {
	StoreId      string
	StoreName    string
	ProductId    string
	ProductName  string
	ProductPrice float64
	Quantity     int
}

type Basket struct {
	ddd.Aggregate
	CustomerId string
	PaymentId  string
	Items      map[string]Item
	Status     BasketStatus
}

func NewBasket(id string) *Basket {
	return &Basket{
		Aggregate: *ddd.NewAggregate(id, BasketAggregate),
		Items:     make(map[string]Item),
	}
}

func StartBasket(id string, customerID string) (*Basket, error) {
	if id == "" {
		return nil, ErrBasketIDCannotBeBlank
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	basket := NewBasket(id)

	basket.AddEvent(BasketStartedEvent, &BasketStarted{
		CustomerID: customerID,
	})

	return basket, nil
}

func (b Basket) IsCancellable() bool {
	return b.Status == BasketIsOpen
}

func (b Basket) IsOpen() bool {
	return b.Status == BasketIsOpen
}

func (b *Basket) Cancel() error {
	if !b.IsCancellable() {
		return ErrBasketCannotBeCancelled
	}

	b.AddEvent(BasketCanceledEvent, &BasketCanceled{})

	return nil
}

func (b *Basket) Checkout(paymentID string) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if len(b.Items) == 0 {
		return ErrBasketHasNoItems
	}

	if paymentID == "" {
		return ErrPaymentIDCannotBeBlank
	}

	b.AddEvent(BasketCheckedOutEvent, &BasketCheckedOut{
		PaymentID:  paymentID,
		CustomerID: b.CustomerId,
		Items:      b.Items,
	})

	return nil
}

func (b *Basket) AddItem(storeId, storeName string,
	productId, productName string, price float64, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	b.AddEvent(BasketItemAddedEvent, &BasketItemAdded{
		Item: Item{
			StoreId:      storeId,
			ProductId:    productId,
			StoreName:    storeName,
			ProductName:  productName,
			ProductPrice: price,
			Quantity:     quantity,
		},
	})

	return nil
}

func (b *Basket) RemoveItem(productId string, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	if _, exists := b.Items[productId]; exists {
		b.AddEvent(BasketItemRemovedEvent, &BasketItemRemoved{
			ProductID: productId,
			Quantity:  quantity,
		})
	}

	return nil
}
