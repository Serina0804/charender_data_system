package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RegisterHandler)

	http.Handle("/", r)

	fmt.Println("boot server")
	http.ListenAndServe(":8080", nil)
}
