package models

import (
	"fmt"
	"log"
	"strconv"
)

type RestaurantLocation struct {
	id            int    `json:"id"`
	streetAddress string `json:"streetAddress"`
	county        string `json:"county"`
	created       int    `json:"created"`
	updated       int    `json:"updated"`
}

type RestaurantLocations []RestaurantLocation

func IsRestaurantlocationIdValid(restaurantLocationId int) bool {
	db, e := myDb.setup()
	defer db.Close()

	if e != nil {
		fmt.Print(e)
	}

	var id int

	selectStatement := "SELECT rl.id "
	selectStatement += "FROM restaurantLocations AS rl "
	selectStatement += "WHERE rl.id = ? "

	err := db.QueryRow(selectStatement, strconv.Itoa(restaurantLocationId)).Scan(&id)
	if err != nil {
		fmt.Print(err)
		return false
	}
	if id > 0 {
		return true
	}
	return false
}

func GetAllRestaurantLocations() RestaurantLocations {
	log.Println("GetAllRestaurantLocations")
	db, e := myDb.setup()
	defer db.Close()

	if e != nil {
		fmt.Print(e)
	}

	selectStatement := "SELECT rl.id, rl.streetAddress, rl.county "
	selectStatement += "FROM restaurantLocations as rl"

	log.Println(selectStatement)

	rows, err := db.Query(selectStatement)
	if err != nil {
		fmt.Print(err)
	}

	var results = make([]RestaurantLocation, 0)

	i := 0
	for rows.Next() {

		var (
			id            int
			streetAddress string
			county        string
		)
		var err = rows.Scan(&id, &streetAddress, &county)

		//		layout := "2006-01-02"

		if err != nil {
			fmt.Println(err)
		}
		transaction := RestaurantLocation{id: id, streetAddress: streetAddress, county: county}
		results = append(results, transaction)
		i++
	}

	return results
}

func (rl *RestaurantLocation) GetId() int {
	return rl.id
}
