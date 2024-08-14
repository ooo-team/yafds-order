package order

import (
	"context"
	"time"

	model "github.com/ooo-team/yafds-order/internal/model/order"
)

func (s *Service) Get(ctx context.Context, orderID uint32) (*model.Order, error) {
	return &model.Order{ID: 0, CustomerID: 0, Status: 0, Timestamp: time.Now()}, nil
}
