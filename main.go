package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "lets gooooo %s", time.Now())
}

func buzz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "buzzzzzzzz")
}

var port = 8080

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/buzz", buzz)
	fmt.Printf("The server is up n running on %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
