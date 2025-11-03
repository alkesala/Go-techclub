package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "lets gooooo %s", time.Now())
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := Response{
		Status: "OK",
	}
	_ = json.NewEncoder(w).Encode(response)

}

func buzz(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "buzzzzzzzz")
}

var port = 8080

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/health", health)
	http.HandleFunc("/buzz", buzz)
	fmt.Printf("The server is up n running on %d\n", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
