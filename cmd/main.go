package main

import (
	"log"
	"net/http"

	"github.com/LokeshvarKr/bookstore/pkg"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	pkg.RegisterRouters(r)
	log.Println("starting server")
	err := http.ListenAndServe(":9005", r)
	if err != nil {
		log.Fatalln("Error in starting server")
	}
}
