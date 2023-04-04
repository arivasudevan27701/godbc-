package main

import (
	"context"
	"fmt"

	con "test/functions"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	client := con.Connector()
	coll := client.Database("test").Collection("restaurants")
	filter := bson.D{{Key: "cuisine", Value: "American "}}
	var result Restaurant
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			fmt.Println("prog ends")
			return
		}
		panic(err)
	}
	fmt.Println(result)
}

type Restaurant struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string
	RestaurantId string `bson:"restaurant_id"`
	Cuisine      string
	Address      interface{}
	Borough      string
	Grades       []interface{}
}
