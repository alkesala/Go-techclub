package main

import (
	"fmt"
	"net/http"
)

var port = 8080 // common development port for Go
func main() {
	fmt.Printf("The server is running on: %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
