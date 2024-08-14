package order

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	model "github.com/ooo-team/yafds-order/internal/model/order"
)

func (s *Service) Create(ctx context.Context, instance *model.Order) (uint32, error) {
	orderUUID, err := uuid.NewUUID() // перенести в отдельную папку service по аналогии с yafds-customer/internal/service
	if err != nil {
		return 0, fmt.Errorf("failed to generate uuid: %v", err)
	}

	log.Printf("Generated uuid.ID = %d", orderUUID.ID())
	err = s.Repo.Create(ctx, orderUUID.ID(), instance.CustomerID)

	return uint32(orderUUID.ID()), err
}
