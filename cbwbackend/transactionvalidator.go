package cbwbackend

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routingfun() {
	router := mux.NewRouter()
	router.HandleFunc("/Insertion", Validator)
	router.HandleFunc("/Update", Updator)
	http.ListenAndServe(":4000", router)

}

type Result struct {
	id     int
	Method string
	Params struct {
		Channel  string
		Refno    string
		Creditor struct {
			Identification     string
			Identificationtype string
		}
		Api struct {
			Credential string
			Apikey     string
			Sig        string
		}
	}
}

type Transaction struct {
	Method string `json:"method,omitempty"`
	Params struct {
		Payload struct {
			Channel string `json:"channel,omitempty"`
			Typec   string `json:"transactionType,omitempty"`
			Refno   string `json:"reference,omitempty"`
			Amnt    struct {
				Amount string `json:"amount,omitempty"`
			} `json:"transactionAmount,omitempty"`
			Creditor struct {
				Identification     string `json:"identification,omitempty"`
				Identificationtype string `json:"identificationtype,omitempty"`
			} `json:"creditorAccount,omitempty"`
		} `json:"payload,omitempty"`
		Api struct {
			Credential string `json:"credential,omitempty"`
			Apikey     string `json:"apiKey,omitempty"`
			Sig        string `json:"signature,omitempty"`
		} `json:"api,omitempty"`
	} `json:"params,omitempty"`
}
