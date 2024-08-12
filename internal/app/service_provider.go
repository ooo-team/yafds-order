package app

import (
	"github.com/ooo-team/yafds-order/internal/repository"
	repositoryOrder "github.com/ooo-team/yafds-order/internal/repository/order"
)

// "github.com/ooo-team/yafds-customer/internal/service"
// customerService "github.com/ooo-team/yafds-customer/internal/service/customer"
// "github.com/ooo-team/yafds-order/internal/repository"
// customerRepository "github.com/ooo-team/yafds-order/internal/repository/customer"

type serviceProvider struct {
	// customerService service.CustomerService
	orderRepo repository.OrderRepository
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

// func (s *serviceProvider) CustomerService() service.CustomerService {
// 	if s.customerService == nil {
// 		s.customerService = customerService.NewService(s.CustomerRepo())
// 	}
// 	return s.customerService
// }
