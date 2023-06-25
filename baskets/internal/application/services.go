package application

import (
	"context"
)

type StoreService interface {
	Find(ctx context.Context, storeID string) (*Store, error)
}

type ProductService interface {
	Find(ctx context.Context, productID string) (*Product, error)
}
