package models

import (
	"fmt"
	"log"
)

type Payment struct {
	id                   int    `json:"id"`
	amount               int    `json:"amount"`
	tableNo              int    `json:"tableNo"`
	restaurantLocationId int    `json:"restaurantLocationId"`
	paymentReferenceId   int    `json:"paymentReferenceId"`
	cardType             string `json:"cardType"`
	created              int    `json:"created"`
	updated              int    `json:"updated"`
}

func addPayment(amount int, tableNo int, paymentReferenceId int) Payment {
	payment := Payment{amount: amount, tableNo: tableNo, paymentReferenceId: paymentReferenceId}
	return payment
}

func (p *Payment) validateRestaurantLocation(restaurantLocationId int) bool {
	return isRestaurantlocationIdValid(restaurantLocationId)
}

func (p *Payment) validateCardType(cardType string) bool {
	if cardType == "visa" || cardType == "mastercard" || cardType == "americanExpress" {
		return true
	}
	return false
}

func (p *Payment) savePayment() {
	db, e := myDb.setup()
	defer db.Close()

	if e != nil {
		fmt.Print(e)
	}

	stmt, err := db.Prepare("INSERT INTO payments (amount, cardType, paymentReferenceId, restaurantLocationId, updated, created) values (?, (SELECT id FROM accountTypes WHERE type = ?), ?, ?, ?, UNIX_TIMESTAMP(), UNIX_TIMESTAMP())")
	if err != nil {
		fmt.Print(err)
		log.Println(err)
	}
	res, err := stmt.Exec(p.amount, p.cardType, p.paymentReferenceId, p.restaurantLocationId)
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
