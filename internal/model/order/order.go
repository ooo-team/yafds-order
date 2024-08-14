package model

import "time"

type OrderState int

const (
	CUSTOMER_CREATED    OrderState = iota
	CUSTOMER_PAID                  = iota
	CUSTOMER_CANCELED              = iota
	KITCHEN_ACCEPTED               = iota
	KITCHEN_PREPARING              = iota
	KITCHEN_DENIED                 = iota
	KITCHEN_REFUNDED               = iota
	DELIVERY_PENDING               = iota
	DELIVERY_PICKING               = iota
	DELIVERY_DELIVERING            = iota
	DELIVERY_COMPLETE              = iota
	DELIVERY_DENIED                = iota
	DELIVERY_REFUNDED              = iota
)

// String - Creating common behavior - give the type a String function
func (s OrderState) String() string {
	return [...]string{
		"CUSTOMER_CREATED",
		"CUSTOMER_PAID",
		"CUSTOMER_CANCELED",
		"KITCHEN_ACCEPTED",
		"KITCHEN_PREPARING",
		"KITCHEN_DENIED",
		"KITCHEN_REFUNDED",
		"DELIVERY_PENDING",
		"DELIVERY_PICKING",
		"DELIVERY_DELIVERING",
		"DELIVERY_COMPLETE",
		"DELIVERY_DENIED",
		"DELIVERY_REFUNDED"}[s]
}

type Order struct {
	ID         uint32
	CustomerID uint32
	Status     uint16
	Timestamp  time.Time
}
