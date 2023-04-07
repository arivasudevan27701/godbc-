package cbwbackend

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	db "test/dbfunctions"

	"go.mongodb.org/mongo-driver/bson"
)

func Updator(w http.ResponseWriter, r *http.Request) {

	byte, _ := io.ReadAll(r.Body)
	k := string(byte)
	p := strings.Split(k, string(rune(34)))
	s, t := sep(p)
	condition := bson.M{"params.api.sig": 1}
	upd := bson.D{{Key: "$set", Value: bson.D{{Key: s, Value: t}}}}
	client := db.Connector()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(k)
	if err != nil {
		log.Println(err)
	}
	db.Update(client, condition, upd, "test", "space")
}
func sep(s []string) (string, string) {
	return s[1], s[len(s)-2]

}
