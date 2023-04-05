package functions

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Find(client *mongo.Client, ref string, Dbname string, Collname string) (bool, error) {
	var r Ref
	coll := client.Database(Dbname).Collection(Collname)
	filter := bson.D{{Key: "params.payload.refno", Value: ref}}
	err := coll.FindOne(context.TODO(), filter).Decode(&r)
	fmt.Println(r)
	if err != nil {
		fmt.Println(r.Params.Payload.Refno)
		return false, err
	}
	if r.Params.Payload.Refno == ref {
		fmt.Println(ref, r.Params.Payload.Refno)
		return true, nil
	}
	return false, nil
}

type Ref struct {
	Params struct {
		Payload struct {
			Refno string `json:"reference"`
		} `json:"payload"`
	} `json:"params"`
}
