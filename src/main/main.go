package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Started . .")
	// its gonna create(if not exist) a table for us, then kick off our RESTAPI
	InitialMigration()
	handleRequest()
}
func handleRequest() {
	θ := mux.NewRouter().StrictSlash(true)
	θ.HandleFunc("/", hello).Methods("GET")

	θ.HandleFunc("/users", AllUsers).Methods("GET")
	θ.HandleFunc("/user/{firstname}/{lastname}/{email}", NewUser).Methods("POST")
	θ.HandleFunc("/user/{firstname}", DeleteUser).Methods("DELETE")
	θ.HandleFunc("/user/{firstname}/{lastname}/{email}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", θ))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Universe!")
}
