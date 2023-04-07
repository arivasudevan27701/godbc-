package functions

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Find(client *mongo.Client, ref string, Dbname string, Collname string) (bool, error) {
	var r Ref
	coll := client.Database(Dbname).Collection(Collname)
	filter :=bson.M{"params.payload.refno":ref}
	err := coll.FindOne(context.TODO(), filter).Decode(&r)
	if err != nil {
		return false, err
	}
	if r.Params.Payload.Refno == ref {
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
