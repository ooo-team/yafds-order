package service

import (
	"context"

	model "github.com/ooo-team/yafds-order/internal/model/order"
)

type OrderService interface {
	Create(ctx context.Context, instance *model.Order) (uint32, error)
	Get(ctx context.Context, orderID uint32) (*model.Order, error)
}
