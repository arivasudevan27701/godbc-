package functions

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func Insertion(client *mongo.Client, b []interface{}, Dbname string, Collname string) {
	Collection := client.Database(Dbname).Collection(Collname)
	insertResult, err := Collection.InsertMany(context.TODO(), b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult)
}
