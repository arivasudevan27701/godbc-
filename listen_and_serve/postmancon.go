package listenandserve

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	con "test/functions"

	"github.com/gorilla/mux"
)

func Mugfunc(w http.ResponseWriter, r *http.Request) {
	byt, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ReadAll Error:", err)
	}
	log.Println("Body :", string(byt))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Value Recived")

}
func Mugfunc2(w http.ResponseWriter, r *http.Request) {
	byt, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ReadAll Error:", err)
	}
	var p interface{}
	log.Println("Body :", string(byt))
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(byt, &p)
	client := con.Connector()
	var i []interface{}
	i = append(i, p)
	q, err := con.Insertion(client, i, "test", "postman")
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(q)
	}
	json.NewEncoder(w).Encode(p)

}

func Main() {
	router := mux.NewRouter()
	router.HandleFunc("/mug", Mugfunc)
	router.HandleFunc("/mux", Mugfunc)
	router.HandleFunc("/mug2", Mugfunc2)
	http.ListenAndServe(":4000", router)

}

//	type perumal struct {
//		Name   string   `json:"name"`
//		Gender string   `json:"gender"`
//		Addres string   `json:"address"`
//		Add2   address2 `json:"address2"`
//	}
//
//	type address2 struct {
//		Houseno int    `json:"houseno"`
//		Street  string `json:"street"`
//		Taluk   string `json:"taluk"`
//	}
