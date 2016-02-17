package main

import (
	"bytes"
	"encoding/json"
	"github.com/vancouverwill/restaurantPaymentsAPI/controllers"
	"github.com/vancouverwill/restaurantPaymentsAPI/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIPOST(t *testing.T) {
	restaurantLocations := models.GetAllRestaurantLocations()
	//	t.Log(restaurantLocations)

	if models.IsRestaurantlocationIdValid(restaurantLocations[0].GetId()) == false {
		t.Error("There are no restaurant locations in the db. please check schema")
	}

	var payment models.Payment
	payment.Amount = 50
	payment.TipAmount = 10
	payment.CardType = "visa"
	payment.RestaurantLocationId = restaurantLocations[0].GetId()
	payment.PaymentReferenceId = 50
	data, err := json.Marshal(payment)
	if err != nil {
		t.Error("error:", err)
	}
	req, err := http.NewRequest("GET", "http://localhost:4000/payments/", bytes.NewBufferString(string(data)))
	if err != nil {
		t.Error("GET did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	controllers.PaymentCreate(w, req)

	if w.Code != 200 && w.Code != 201 {
		t.Error("GET did not work as expected. the status was not ", http.StatusOK, ", it was ", w.Code)
	}

	t.Log("status:", w.Code, "body:", w.Body.String())
}
