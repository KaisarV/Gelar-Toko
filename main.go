package main

import (
	controller "GelarToko/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", controller.UserLogin).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := corsHandler.Handler(router)
	log.Println("Starting on Port")

	err := http.ListenAndServe(":8080", handler)
	log.Fatal(err)

}
