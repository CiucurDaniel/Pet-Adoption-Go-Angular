package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server ...")

	router := mux.NewRouter()
	router.HandleFunc("/", Hello)

	// TODO: on root give the angular frontend
	// http.Handle("/", ...angular-frontend)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello from server!")
}