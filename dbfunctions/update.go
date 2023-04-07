package functions

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Update(client *mongo.Client, condition primitive.M, update primitive.D, db string, coll string) error {
	collection := client.Database(db).Collection(coll)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.UpdateOne(ctx, condition, update)

	if err != nil {
		fmt.Println(err)
	}
	return nil
}
