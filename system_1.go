package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", loginHandler)
	http.Handle("/", r)

	fmt.Println("boot server")
	http.ListenAndServe(":8080", nil)
}
