package controllers

import (
	"encoding/json"
	//	"fmt"
	"github.com/vancouverwill/restaurantPaymentsAPI/models"
	//	"log"
	"net/http"
	//	"strconv"
	"io"
	"io/ioutil"
)

func PaymentsIndex(response http.ResponseWriter, request *http.Request) {

}

func PaymentsByRestaurant(response http.ResponseWriter, request *http.Request) {

}

func PaymentCreate(response http.ResponseWriter, request *http.Request) {
	var payment models.Payment
	payment = jsonToObject(response, request, payment)

	Payment

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(response).Encode(transaction); err != nil {
		panic(err)
	}
}

func jsonToObject(response http.ResponseWriter, request *http.Request, payment models.Payment) models.Payment {
	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := request.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &payment); err != nil {
		response.Header().Set("Content-Type", "application/json; charset=UTF-8")
		response.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(response).Encode(err); err != nil {
			panic(err)
		}
	}
	return payment
}
