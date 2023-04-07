package cbwbackend

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	con "test/dbfunctions"

	"go.mongodb.org/mongo-driver/bson"
)

func Validator(w http.ResponseWriter, r *http.Request) {
	res := Result{}
	b := true
	byt, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ReadAll Error:", err)
	}
	var p Transaction

	// log.Println("Body :", string(byt))
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(byt, &p)
	client := con.Connector()
	defer client.Disconnect(context.TODO())

	var i []interface{}
	i = append(i, p)

	if b && p.Params.Api.Credential != "Basic YXJpdmFzdWRldmFuLnMrNkBuZXR4ZC5jb206NzEyMGY0NWFmM2ZjNGIxMGJhOGUzYjY2YTQ5YzBlMGU==" {
		json.NewEncoder(w).Encode("invalid credential")
		res.Params.Api.Credential = "error in validatiing credenetial"
		b = false

	}
	if b && p.Params.Api.Apikey != "7120f45af3fjh45h72h6f8rg0h4j3e" {
		json.NewEncoder(w).Encode("error in api key !!!!")
		res.Params.Api.Apikey = "error in validatiing credenetial API KEY"
		b = false

	}
	if b && p.Params.Payload.Channel != "ACH" {
		json.NewEncoder(w).Encode("Invalid Transaction")
		res.Params.Channel = "Channel Not Found"
		b = false
	}
	if b && p.Params.Payload.Creditor.Identification != "200542309118191" {
		json.NewEncoder(w).Encode("error in finding the account")
		res.Params.Creditor.Identification = "creditor idedntification not found"
		b = false
	}
	ref := p.Params.Payload.Refno
	b2, err2 := con.Find(client, ref, "test", "postman")
	if err2 != nil {
		json.NewEncoder(w).Encode(err2)
	}
	if b && b2 {
		json.NewEncoder(w).Encode("Duplicate ref no")
		res.Params.Refno = "duplicate ref no"
		b = false
	}

	if b {
		res.id = 10001
		json.NewEncoder(w).Encode(p)
		_, err = con.Insertion(client, i, "test", "postman")
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
	} else {
		json.NewEncoder(w).Encode(res)
	}
	condition := bson.M{"method": "value change"}
	upd := bson.D{{Key: "$set", Value: bson.D{{Key: "method", Value: "again changed"}}}}

	con.Update(client, condition, upd, "test", "postman")

}
