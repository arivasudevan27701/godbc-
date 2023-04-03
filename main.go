package main

import (
	con "test/functions"
)

func main() {
	client := con.Connector()
	docs := []interface{}{
		con.Tea{Type: "Masala", Rating: 10, Vendor: []string{"A", "C"}},
		con.Tea{Type: "English Breakfast", Rating: 6},
		con.Tea{Type: "Oolong", Rating: 7, Vendor: []string{"C"}},
		con.Tea{Type: "Assam", Rating: 5},
		con.Tea{Type: "Earl Grey", Rating: 8, Vendor: []string{"A", "B"}},
	}
	con.Insertion(client, docs, "test", "test")

}
