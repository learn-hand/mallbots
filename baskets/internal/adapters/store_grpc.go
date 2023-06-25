package adapters

import (
	"context"

	"github.com/learn-hand/mallbots/baskets/internal/application"
	"github.com/learn-hand/mallbots/stores/storespb"
	"google.golang.org/grpc"
)

type StoreGrpc struct {
	client storespb.StoresServiceClient
}

var _ application.StoreService = (*StoreGrpc)(nil)

func NewStoreGrpc(conn *grpc.ClientConn) StoreGrpc {
	return StoreGrpc{client: storespb.NewStoresServiceClient(conn)}
}

func (r StoreGrpc) Find(ctx context.Context, storeID string) (*application.Store, error) {
	resp, err := r.client.GetStore(ctx, &storespb.GetStoreRequest{
		Id: storeID,
	})
	if err != nil {
		return nil, err
	}

	return r.storeToDomain(resp.Store), nil
}

func (r StoreGrpc) storeToDomain(store *storespb.Store) *application.Store {
	return &application.Store{
		ID:       store.GetId(),
		Name:     store.GetName(),
		Location: store.GetLocation(),
	}
}
