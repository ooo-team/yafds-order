package repository

import (
	"context"

	model "github.com/ooo-team/yafds-order/internal/model/order"
)

type OrderRepository interface {
	Create(ctx context.Context, orderID uint32, customerID uint32) error
	Get(ctx context.Context, orderID uint32) (*model.Order, error)
}
