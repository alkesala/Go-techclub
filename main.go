package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	resp := Response{
		Status: "OK",
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("InternalserverError %v", err)
		http.Error(w, "Internalserver error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

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
