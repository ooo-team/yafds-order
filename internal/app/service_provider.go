package app

import (
	"github.com/ooo-team/yafds-order/internal/repository"
	repositoryOrder "github.com/ooo-team/yafds-order/internal/repository/order"
	"github.com/ooo-team/yafds-order/internal/service"
	orderService "github.com/ooo-team/yafds-order/internal/service/order"
)

type serviceProvider struct {
	customerService service.OrderService
	orderRepo       repository.OrderRepository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) OrderRepo() repository.OrderRepository {
	if s.orderRepo == nil {
		s.orderRepo = repositoryOrder.NewOrderRepository()
	}
	return s.orderRepo
}

func (s *serviceProvider) OrderService() service.OrderService {
	if s.customerService == nil {
		s.customerService = &orderService.Service{Repo: s.OrderRepo()}
	}
	return s.customerService
}
