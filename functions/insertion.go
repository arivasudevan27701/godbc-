package functions

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func Insertion(client *mongo.Client, b []interface{}, Dbname string, Collname string) (*mongo.InsertManyResult, error) {
	Collection := client.Database(Dbname).Collection(Collname)
	insertResult, err := Collection.InsertMany(context.TODO(), b)
	if err != nil {
		return nil, err
	}
	fmt.Println("Inserted a single document: ", insertResult)
	return insertResult, nil
}

// func createProfile(w http.ResponseWriter, r *http.Request) {
// 	var userCollection = Client.Database("goTest").Collection("users")
// 	w.Header().Set("Content-Type", "application/json") // for adding       //Content-type
// 	var person user
// 	err := json.NewDecoder(r.Body).Decode(&person) // storing in person   //variable of type user
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	insertResult, err := userCollection.InsertOne(context.TODO(), person)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	json.NewEncoder(w).Encode(insertResult.InsertedID) // return the //mongodb ID of generated document
// }

// type user struct {
// 	Name string `json:"name"`
// 	City string `json:"city"`
// 	Age  int    `json:"age"`
// }
