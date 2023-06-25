package monolith

import (
	"context"
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Module interface {
	Startup(context.Context, Monolith) error
}

type Monolith struct {
	cfg    AppConfig
	db     *sql.DB
	logger zerolog.Logger
	mux    *chi.Mux
	rpc    *grpc.Server

	modules []Module
}

func (m Monolith) Config() AppConfig      { return m.cfg }
func (m Monolith) DB() *sql.DB            { return m.db }
func (m Monolith) Logger() zerolog.Logger { return m.logger }
func (m Monolith) Mux() *chi.Mux          { return m.mux }
func (m Monolith) Rpc() *grpc.Server      { return m.rpc }

func (m *Monolith) startupModules() error {
	for _, module := range m.modules {
		if err := module.Startup(context.TODO(), *m); err != nil {
			return err
		}
	}
	return nil
}
