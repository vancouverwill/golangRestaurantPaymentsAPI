package models

import (
	"fmt"
	"log"
)

type Payment struct {
	id                   int    `json:"id"`
	Amount               int    `json:"amount"`
	TipAmount            int    `json:"tipAmount"`
	TableNo              int    `json:"tableNo"`
	RestaurantLocationId int    `json:"restaurantLocationId"`
	PaymentReferenceId   int    `json:"paymentReferenceId"`
	CardType             string `json:"cardType"`
	created              int    `json:"created"`
	updated              int    `json:"updated"`
}

func AddPayment(amount int, tableNo int, paymentReferenceId int) Payment {
	payment := Payment{Amount: amount, TableNo: tableNo, PaymentReferenceId: paymentReferenceId}
	return payment
}

func (p *Payment) IsValidRestaurantLocation() bool {
	return IsRestaurantlocationIdValid(p.RestaurantLocationId)
}

func (p *Payment) IsValidCardType() bool {
	if p.CardType == "visa" || p.CardType == "mastercard" || p.CardType == "americanExpress" {
		return true
	}
	return false
}

func (p *Payment) SavePayment() {
	db, e := myDb.setup()
	defer db.Close()

	if e != nil {
		fmt.Print(e)
	}

	stmt, err := db.Prepare("INSERT INTO payments (amount, tipAmount, cardType, paymentReferenceId, restaurantLocationId) values (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Print(err)
		log.Println(err)
	}
	res, err := stmt.Exec(p.Amount, p.TipAmount, p.CardType, p.PaymentReferenceId, p.RestaurantLocationId)
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	RowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("RowsAffected", RowsAffected)
	log.Println("transaction entered", lastId)
}

func (p *Payment) GetAmount() int {
	return p.Amount
}
