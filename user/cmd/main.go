package main

import (
	"log"
	"net/http"
	"user/internal/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("", handler.UserRegistration())
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// func UserRegistration(w http.ResponseWriter, r *http.Request) {

// }
