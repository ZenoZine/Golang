package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var members = map[string]string{
	"1": "Andy",
	"2": "Peter",
	"3": "Gabriella",
	"4": "Jordy",
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(members)
}

func deleteMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, ok := members[id]; ok {
		delete(members, id)
		json.NewEncoder(w).Encode(members)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(members)
	}
}

func main() {
	// Used a different method in the solution:
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/members", getMembers).Methods("GET")
	router.HandleFunc("/deleteMembers/{id}", deleteMembers).Methods("DELETE")

	fmt.Println("Server starting on port 3000")
	http.ListenAndServe(":3000", router)
}
