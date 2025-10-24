package main

import (
	"log"
	"net/http"
	"github.com/homirock/ecommerce-backend-go/user/internal/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("", handler.UserRegistration)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
