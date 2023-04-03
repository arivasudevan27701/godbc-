package functions

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Publisher string
	Copies    int
}
type Objects struct {
	Name   string
	Age    int
	Degree string
	Skill  Skills
}
type Skills struct {
	Primary   []string
	Secondary []string
	Add       Address
}
type Address struct {
	Street   string
	Door_num int
	District string
}
type Tea struct {
	Type   string
	Rating int32
	Vendor []string `bson:"vendor,omitempty" json:"vendor,omitempty"`
}

func (o Objects) Struct_Insertion(client *mongo.Client, Dbname string, Collname string) {
	Collection := client.Database(Dbname).Collection(Collname)
	insertResult, err := Collection.InsertOne(context.TODO(), o)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
