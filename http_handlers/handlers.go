package http_handlers

import (
	"fmt"
	"net/http"
)

func GetCars(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Tesla and Lambo are cars")
}
