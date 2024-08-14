package order

import (
	"github.com/ooo-team/yafds-order/internal/repository"
	def "github.com/ooo-team/yafds-order/internal/service"
)

type Service struct {
	Repo repository.OrderRepository
}

var _ def.OrderService = (*Service)(nil)
