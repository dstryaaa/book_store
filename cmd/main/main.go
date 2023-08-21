package main

import (
	"fmt"
	"log"
	"main/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server at the 8080 port...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
