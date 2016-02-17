package main

import (
	"github.com/vancouverwill/restaurantPaymentsAPI/controllers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"PaymentsIndex",
		"GET",
		"/transactions",
		controllers.PaymentsIndex,
	},
	Route{
		"PaymentsByRestaurant",
		"GET",
		"/payments/?restaurantId={restaurantLocationId}",
		controllers.PaymentsSearch,
	},
	Route{
		"PaymentCreate",
		"POST",
		"/payments",
		controllers.PaymentCreate,
	},
}
