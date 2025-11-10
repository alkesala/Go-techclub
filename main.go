package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// omitempty == optional

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := Response{
		Status: "OK",
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func brrr(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Brrrrrrrrrr %s", time.Now())
}

func buzz(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "There is alot buzzwords in the IT sector")
}

var port = 8080 // common development port for Go
func main() {

	http.HandleFunc("/brrr", brrr)
	http.HandleFunc("/buzz", buzz)
	http.HandleFunc("/health", health)
	fmt.Printf("The server is running on: %d\n", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
