package models

type RestaurantLocation struct {
	id            int    `json:"id"`
	streetAddress string `json:"streetAddress"`
	county        string `json:"county"`
	created       int    `json:"created"`
	updated       int    `json:"updated"`
}


isRestaurantlocationIdValid(restaurantLocationId int) bool {
	
}