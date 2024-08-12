package model

import "time"

type Order struct {
	ID         uint32
	CustomerID uint32
	Status     uint16
	Timestamp  time.Time
}
