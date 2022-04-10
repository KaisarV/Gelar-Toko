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

	//1. User Biasa 2. Memiliki Toko 3. Admin

	router.HandleFunc("/login", controller.UserLogin).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")
	router.HandleFunc("/register", controller.InsertUser).Methods("POST")

	router.HandleFunc("/users", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controller.Authenticate(controller.DeleteUser, 3)).Methods("DELETE")
	router.HandleFunc("/users", controller.Authenticate(controller.UpdateMyProfile, 1)).Methods("PUT")

	router.HandleFunc("/stores", controller.GetAllStores).Methods("GET")
	router.HandleFunc("/stores", controller.Authenticate(controller.InsertMyStore, 1)).Methods("POST")
	router.HandleFunc("/stores", controller.Authenticate(controller.DeleteMyStore, 2)).Methods("DELETE")
	router.HandleFunc("/stores", controller.Authenticate(controller.UpdateMyStore, 2)).Methods("PUT")

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
