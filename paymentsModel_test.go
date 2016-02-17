package main

import (
	"github.com/vancouverwill/restaurantPaymentsAPI/models"
	"testing"
)

func TestAddPayment(t *testing.T) {
	t.Log("TestAddPayment")
	payment := models.addPayment(1000, 1000, 10000)

	if payment.amount != 1000 {
		t.Error("TestAddPayment() did not work as expected. The amount did not equal ", 1000)
	}
	t.Log("TestAddPayment() successful")
}
