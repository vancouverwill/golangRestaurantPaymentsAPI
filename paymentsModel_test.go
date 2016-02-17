package main

import (
	"github.com/vancouverwill/restaurantPaymentsAPI/models"
	"testing"
)

func TestAddPayment(t *testing.T) {
	t.Log("TestAddPayment")
	payment := models.AddPayment(1000, 1000, 10000)

	if payment.GetAmount() != 1000 {
		t.Error("TestAddPayment() did not work as expected. The amount did not equal ", 1000)
	}
	t.Log("TestAddPayment() successful")
}
