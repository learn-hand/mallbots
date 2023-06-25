package adapters

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/learn-hand/mallbots/baskets/internal/domain"
	"github.com/learn-hand/mallbots/internal/ddd"
	"github.com/stackus/errors"
)

type PostgreBasketRepository struct {
	tableName string
	db        *sql.DB
}

func NewPostgreBasketRepository(tableName string, db *sql.DB) PostgreBasketRepository {
	return PostgreBasketRepository{tableName: tableName, db: db}
}

func (r PostgreBasketRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}

func (r PostgreBasketRepository) Find(ctx context.Context, basketID string) (*domain.Basket, error) {
	const query = "SELECT customer_id, payment_id, items, status FROM %s WHERE id = $1 LIMIT 1"

	basket := &domain.Basket{
		Aggregate: *ddd.NewAggregate(basketID, domain.BasketAggregate),
	}
	var items []byte
	var status string

	err := r.db.QueryRowContext(ctx, r.table(query), basketID).Scan(&basket.CustomerId, &basket.PaymentId, &items, &status)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	basket.Status, err = domain.FromString(status)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(items, &basket.Items)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	return basket, nil
}

func (r PostgreBasketRepository) Save(ctx context.Context, basket *domain.Basket) error {
	const query = "INSERT INTO %s (id, customer_id, payment_id, items, status) VALUES ($1, $2, $3, $4, $5)"

	items, err := json.Marshal(basket.Items)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}

	_, err = r.db.ExecContext(ctx, r.table(query),
		basket.AggregateId(), basket.CustomerId, basket.PaymentId, items, basket.Status.String())

	return errors.ErrInternalServerError.Err(err)
}

func (r PostgreBasketRepository) Update(ctx context.Context, basket *domain.Basket) error {
	const query = "UPDATE %s SET customer_id = $2, payment_id = $3, items = $4, status = $5  WHERE id = $1"

	items, err := json.Marshal(basket.Items)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}

	_, err = r.db.ExecContext(ctx, r.table(query),
		basket.AggregateId(), basket.CustomerId, basket.PaymentId, items, basket.Status.String())

	return errors.ErrInternalServerError.Err(err)
}

func (r PostgreBasketRepository) DeleteBasket(ctx context.Context, basketID string) error {
	const query = "DELETE FROM %s WHERE id = $1 LIMIT 1"

	_, err := r.db.ExecContext(ctx, r.table(query), basketID)

	return errors.ErrInternalServerError.Err(err)
}
