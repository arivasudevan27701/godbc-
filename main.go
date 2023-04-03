package main

import (
	con "test/functions"
)

func main() {
	client := con.Connector()
	r := []interface{}{
		Restaurant{
			ID:           901,
			Name:         "dlshjs",
			RestaurantId: "sdkjb",
			Cuisine:      "chinese",
			Address:      []any{"sdalkh", "sajhs", 90},
			Borough:      "kjsabdkbb",
			Grades: []any{
				"afk",
				90,
				"alkdhl",
			},
		},
		Restaurant{
			ID:           902,
			Name:         "dlshjs",
			RestaurantId: "sdkjb",
			Cuisine:      "chinese",
			Address:      []any{"sdalkh", "sajhs", 90},
			Borough:      "kjsabdkbb",
			Grades: []any{
				"afk",
				90,
				"alkdhl",
			},
		},
	}

	con.Insertion(client, r, "test", "test")
}

type Restaurant struct {
	ID           int `bson:"_id"`
	Name         string
	RestaurantId string `bson:"restaurant_id"`
	Cuisine      string
	Address      interface{}
	Borough      string
	Grades       []interface{}
}
