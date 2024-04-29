package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RegisterHandler)
	r.HandleFunc("/index.html", IndexHandler) // 追加

	http.Handle("/", r)

	fmt.Println("boot server")
	http.ListenAndServe(":8080", nil)
}
