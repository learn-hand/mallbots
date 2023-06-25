package adapters

import (
	"context"

	"github.com/learn-hand/mallbots/baskets/internal/application"
	"github.com/learn-hand/mallbots/stores/storespb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type ProductGrpc struct {
	client storespb.StoresServiceClient
}

func NewProductGrpc(conn *grpc.ClientConn) ProductGrpc {
	return ProductGrpc{client: storespb.NewStoresServiceClient(conn)}
}

func (r ProductGrpc) Find(ctx context.Context, productID string) (*application.Product, error) {
	resp, err := r.client.GetProduct(ctx, &storespb.GetProductRequest{
		Id: productID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "requesting product")
	}

	return r.productToDomain(resp.Product), nil
}

func (r ProductGrpc) productToDomain(product *storespb.Product) *application.Product {
	return &application.Product{
		ID:      product.GetId(),
		StoreID: product.GetStoreId(),
		Name:    product.GetName(),
		Price:   product.GetPrice(),
	}
}
