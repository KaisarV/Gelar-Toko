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
	//Verifikasi akun setelah login dengan link yang dikirim ke email
	router.HandleFunc("/verify/{token}", controller.VerifyToken).Methods("GET")
	//Verifikasi akun manual(untuk testing tanpa lewat email(email dummy))
	router.HandleFunc("/verify/testing/{id}", controller.VerifyTokenById).Methods("GET")

	router.HandleFunc("/users/GetAllUsers", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/DeleteUser/{id}", controller.Authenticate(controller.DeleteUser, 3)).Methods("DELETE")
	router.HandleFunc("/users/UpdateProfile", controller.Authenticate(controller.UpdateMyProfile, 1)).Methods("PUT")

	router.HandleFunc("/stores/GetAllStores", controller.GetAllStores).Methods("GET")
	router.HandleFunc("/stores/InsertMyStore", controller.Authenticate(controller.InsertMyStore, 1)).Methods("POST")
	router.HandleFunc("/stores/DeleteMyStore", controller.Authenticate(controller.DeleteMyStore, 2)).Methods("DELETE")
	router.HandleFunc("/stores/UpdateMyStore", controller.Authenticate(controller.UpdateMyStore, 2)).Methods("PUT")

	//transactions
	router.HandleFunc("/transactions/GetTransactions", controller.Authenticate(controller.GetTransaction, 1)).Methods("GET")
	router.HandleFunc("/transactions/InsertNewTransactions", controller.Authenticate(controller.InsertNewTransactions, 1)).Methods("POST")

	//products
	router.HandleFunc("/products", controller.GetAllProduct).Methods("GET")
	router.HandleFunc("/products", controller.Authenticate(controller.InsertNewProduct, 1)).Methods("POST")
	router.HandleFunc("/products/update/{id}", controller.Authenticate(controller.UpdateProduct, 1)).Methods("PUT")

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
