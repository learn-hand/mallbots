package baskets

import (
	"context"

	"github.com/learn-hand/mallbots/baskets/internal/adapters"
	"github.com/learn-hand/mallbots/baskets/internal/application"
	"github.com/learn-hand/mallbots/baskets/internal/ports/rest"
	grpc "github.com/learn-hand/mallbots/baskets/internal/ports/rpc"
	"github.com/learn-hand/mallbots/internal/monolith"
	"github.com/learn-hand/mallbots/internal/rpc"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	// domainDispatcher := ddd.NewEventDispatcher()
	baskets := adapters.NewPostgreBasketRepository("baskets.baskets", mono.DB())
	conn, err := rpc.Dial(ctx, mono.Config().RpcAddress)
	if err != nil {
		return err
	}
	stores := adapters.NewStoreGrpc(conn)
	products := adapters.NewProductGrpc(conn)

	// setup application
	app := application.New(baskets, stores, products)
	// app := logging.LogApplicationAccess(
	// 	application.New(baskets, stores, products, orders, domainDispatcher),
	// 	mono.Logger(),
	// )
	// orderHandlers := logging.LogDomainEventHandlerAccess(
	// 	application.NewOrderHandlers(orders),
	// 	mono.Logger(),
	// )

	// setup ports
	if err := grpc.RegisterServer(app, mono.Rpc()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().RpcAddress); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	// handlers.RegisterOrderHandlers(orderHandlers, domainDispatcher)

	return nil
}
