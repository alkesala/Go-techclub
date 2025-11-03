package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type RoastResponse struct {
	ID    int    `json:"id"`
	Roast string `json:"roast"`
}

var roasts = []RoastResponse{
	{ID: 1, Roast: "You debug code like a blind grandpa"},
}
var nextID = 2

// Handle the json encoding error gracefully. Does anyone spot inconsistency?
func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := Response{
		Status: "OK",
	}
	_ = json.NewEncoder(w).Encode(response)
}

func get(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(roasts)
	if err != nil {
		fmt.Print("Error encoding json", err)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	var createRoast RoastResponse

	err := json.NewDecoder(r.Body).Decode(&createRoast)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(Response{Status: "error", Message: "Invalid JSON"})
		return
	}
	createRoast.ID = nextID
	nextID++
	roasts = append(roasts, createRoast)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(createRoast)
}

func roastHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		get(w, r)
	case "POST":
		create(w, r)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(Response{Status: "error", Message: "Method not allowed"})
	}
}

var port = 8080

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/roasts", roastHandler)
	fmt.Printf("The server is up n running on %d\n", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
