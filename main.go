package main

import (
	controller "GelarToko/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", controller.UserLogin).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)

}
