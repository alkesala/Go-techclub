package main

import (
	"fmt"
	"net/http"
	"time"
)

func brrr(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Brrrrrrrrrr %s", time.Now())
}

var port = 8080 // common development port for Go
func main() {

	http.HandleFunc("/brrr", brrr)
	fmt.Printf("The server is running on: %d\n", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
