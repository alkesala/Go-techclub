package main

import (
	"fmt"
	"net/http"
	"time"
)

// TODO add health endpoint

func brrr(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Brrrrrrrrrr %s", time.Now())
}

func buzz(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "There is alot buzzwords in the IT sector")
}

var port = 8080 // common development port for Go
func main() {

	http.HandleFunc("/brrr", brrr)
	http.HandleFunc("/buzz", buzz)
	//http.HandleFunc("/health", health)
	fmt.Printf("The server is running on: %d\n", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
