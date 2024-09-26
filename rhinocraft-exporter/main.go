package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Concentration struct {
	MaxQuantity int    `json:"maxQuantity"`
	Quantity    int    `json:"quantity"`
	Name        string `json:"name"`
	Time        int    `json:"time"`
}

func main() {
	testData := Concentration{
		MaxQuantity: 1000,
		Quantity:    500,
		Name:        "Alchemy",
		Time:        1727308828,
	}
	response, _ := json.Marshal(testData)

	fmt.Println("init")
	http.HandleFunc("/concentration", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(response))
	})

	log.Fatal(http.ListenAndServe(":8099", nil))

}
