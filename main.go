package main

//importing all req lib

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/laundry/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterLaundryRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
