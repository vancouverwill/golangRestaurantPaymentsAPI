package main

import (
	"github.com/vancouverwill/restaurantPaymentsAPI/models"
	"testing"
)

func TestGetAllRestaurantLocations(t *testing.T) {
	t.Log("TestGetAllRestaurantLocations")
	restaurantLocations := models.GetAllRestaurantLocations()

	if len(restaurantLocations) <= 0 {
		t.Error("TestGetAllRestaurantLocations() did not work as expected. No restaurantLocations locations were returned")
	}
	t.Log("TestAddPayment() successful")
}

func TestIsRestaurantlocationIdValid(t *testing.T) {

	restaurantLocations := models.GetAllRestaurantLocations()
	t.Log(restaurantLocations)

	restaurantLocationIsValid := models.IsRestaurantlocationIdValid(restaurantLocations[0].GetId())

	if restaurantLocationIsValid == false {
		t.Error("TestIsRestaurantlocationIdValid is broken")
	}

}
