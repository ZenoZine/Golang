// Use ^ + C to close the local server

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var dictionary = map[string]string{
	"Go":     "A programming language created by Google",
	"Gopher": "A software engineer who builds with Go",
	"Golang": "Another name for Go",
}

// Returns a JSON version of the dictionary created above
func getDictionary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// In order to return the dictionary as a JSON version, an encoder is required
	json.NewEncoder(w).Encode(dictionary)
}

func create(w http.ResponseWriter, r *http.Request) {
	// 1. Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// 2. Keep track of new entry
	var newEntry map[string]string

	// 3. Read the request
	reqBody, _ := io.ReadAll(r.Body)

	// 4. Parse JSON body
	json.Unmarshal(reqBody, &newEntry)

	// 5. Add new entry to dictionary
	// - Respond with conflict if entry exists
	// - Respond with OK if entry does not already exist

	for k, v := range newEntry {
		if _, ok := dictionary[k]; ok {
			w.WriteHeader(http.StatusConflict)
		} else {
			dictionary[k] = v
			w.WriteHeader(http.StatusCreated)
		}
	}
	// 6. Return dictionary
	json.NewEncoder(w).Encode(dictionary)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/dictionary", getDictionary).Methods("GET")
	router.HandleFunc("/dictionary", create).Methods("POST")

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
