package order

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	commonRepo "github.com/ooo-team/yafds-common/pkg/repository"
	model "github.com/ooo-team/yafds-order/internal/model/order"
)

type OrderRepositoryError struct {
	Message string
}

func (e *OrderRepositoryError) Error() string {
	return e.Message
}

type repository struct {
	db *sql.DB
}

func NewOrderRepository() *repository {
	return &repository{}
}

func (r *repository) GetDB() *sql.DB {

	if r.db != nil {
		return r.db
	}

	r.db = commonRepo.GetDB()

	return r.db
}

func (r *repository) Get(ctx context.Context, orderID uint32) (*model.Order, error) {
	return nil, nil
}

type Result struct {
	Exists bool
	Err    error
}

func (r *repository) checkIfCustomerExists(ctx context.Context, customerID uint32, resultChan chan<- Result) {
	var exists bool
	queryText := `
	select 1
	  from customers c
	 where c.id = $1
	`

	err := r.GetDB().QueryRowContext(ctx, queryText, customerID).Scan(&exists)

	resultChan <- Result{exists, err}

}

func (r *repository) Create(ctx context.Context, orderID uint32, customerID uint32) error {
	resultChan := make(chan Result)
	go r.checkIfCustomerExists(ctx, customerID, resultChan)

	res := <-resultChan

	if res.Err != nil || !res.Exists {
		errMsg := fmt.Sprintf("%s%d", "Заказчик не найден по переданному идентификатору: ", customerID)
		return &OrderRepositoryError{errMsg}
	}

	queryText := `
	insert into orders (id, customer_id, restaurant_id, courier_id, status, timestamp)
	values ($1, $2, 228, 228, $3, $4)
	`

	_, err := r.GetDB().ExecContext(ctx, queryText, orderID, customerID, model.CUSTOMER_CREATED, time.Now())

	return err
}
