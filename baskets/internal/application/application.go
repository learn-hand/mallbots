package application

import (
	"context"

	"github.com/learn-hand/mallbots/baskets/internal/domain"
	"github.com/pkg/errors"
)

type StartBasket struct {
	ID         string
	CustomerID string
}

type CancelBasket struct {
	ID string
}

type CheckoutBasket struct {
	ID        string
	PaymentID string
}

type AddItem struct {
	ID        string
	ProductID string
	Quantity  int
}

type RemoveItem struct {
	ID        string
	ProductID string
	Quantity  int
}

type GetBasket struct {
	ID string
}

type App interface {
	StartBasket(ctx context.Context, start StartBasket) error
	CancelBasket(ctx context.Context, cancel CancelBasket) error
	CheckoutBasket(ctx context.Context, checkout CheckoutBasket) error
	AddItem(ctx context.Context, add AddItem) error
	RemoveItem(ctx context.Context, remove RemoveItem) error
	GetBasket(ctx context.Context, get GetBasket) (*domain.Basket, error)
}

type Application struct {
	baskets  domain.BasketRepository
	stores   StoreService
	products ProductService
}

func New(
	baskets domain.BasketRepository,
	stores StoreService,
	products ProductService) *Application {
	return &Application{
		baskets:  baskets,
		stores:   stores,
		products: products,
	}
}

func (a Application) StartBasket(ctx context.Context, start StartBasket) error {
	basket, err := domain.StartBasket(start.ID, start.CustomerID)
	if err != nil {
		return err
	}

	if err = a.baskets.Save(ctx, basket); err != nil {
		return err
	}

	return nil
}

func (a Application) CancelBasket(ctx context.Context, cancel CancelBasket) error {
	basket, err := a.baskets.Find(ctx, cancel.ID)
	if err != nil {
		return err
	}

	err = basket.Cancel()
	if err != nil {
		return err
	}

	if err = a.baskets.Update(ctx, basket); err != nil {
		return err
	}

	return nil
}

func (a Application) CheckoutBasket(ctx context.Context, checkout CheckoutBasket) error {
	basket, err := a.baskets.Find(ctx, checkout.ID)
	if err != nil {
		return err
	}

	err = basket.Checkout(checkout.PaymentID)
	if err != nil {
		return errors.Wrap(err, "baskets checkout")
	}

	if err = a.baskets.Update(ctx, basket); err != nil {
		return errors.Wrap(err, "basket checkout")
	}

	return nil
}

func (a Application) AddItem(ctx context.Context, add AddItem) error {
	basket, err := a.baskets.Find(ctx, add.ID)
	if err != nil {
		return err
	}

	product, err := a.products.Find(ctx, add.ProductID)
	if err != nil {
		return err
	}

	store, err := a.stores.Find(ctx, product.StoreID)
	if err != nil {
		return nil
	}

	err = basket.AddItem(store.ID, store.Name, product.ID, product.Name, product.Price, add.Quantity)
	if err != nil {
		return err
	}

	if err = a.baskets.Update(ctx, basket); err != nil {
		return err
	}

	return nil
}

func (a Application) RemoveItem(ctx context.Context, remove RemoveItem) error {
	product, err := a.products.Find(ctx, remove.ProductID)
	if err != nil {
		return err
	}

	basket, err := a.baskets.Find(ctx, remove.ID)
	if err != nil {
		return err
	}

	err = basket.RemoveItem(product.ID, remove.Quantity)
	if err != nil {
		return err
	}

	if err = a.baskets.Update(ctx, basket); err != nil {
		return err
	}

	return nil
}

func (a Application) GetBasket(ctx context.Context, get GetBasket) (*domain.Basket, error) {
	return a.baskets.Find(ctx, get.ID)
}
