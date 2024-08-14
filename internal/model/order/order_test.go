package model

import (
	"fmt"
	"testing"
)

func TestOrderStates(t *testing.T) {
	var state OrderState = CUSTOMER_CREATED

	fmt.Println(state)
}
