package main

import (
	"fmt"
	"net/http"
)

var cities = []string{"Tokyo", "Dehli", "Shanghai", "Sao Paulo", "Mexico City"}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world :]\n")
}

func cityList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of most populace cities:\n")

	for _, city := range cities {
		fmt.Fprintf(w, "%s\n", city)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/cityList", cityList)

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", nil)
}
