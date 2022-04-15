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

	router.HandleFunc("/users", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controller.Authenticate(controller.DeleteUser, 3)).Methods("DELETE")
	router.HandleFunc("/users", controller.Authenticate(controller.UpdateMyProfile, 1)).Methods("PUT")

	router.HandleFunc("/stores", controller.GetAllStores).Methods("GET")
	router.HandleFunc("/stores", controller.Authenticate(controller.InsertMyStore, 1)).Methods("POST")
	router.HandleFunc("/stores", controller.Authenticate(controller.DeleteMyStore, 2)).Methods("DELETE")
	router.HandleFunc("/stores", controller.Authenticate(controller.UpdateMyStore, 2)).Methods("PUT")

	//Chat
	router.HandleFunc("/chat/{receiverId}/{userType}", controller.Authenticate(controller.SendChat, 1)).Methods("POST")
	router.HandleFunc("/chat/{userType}", controller.Authenticate(controller.GetChat, 1)).Methods("GET")

	//transactions
	router.HandleFunc("/transactions/GetTransactions", controller.Authenticate(controller.GetTransaction, 1)).Methods("GET")
	router.HandleFunc("/transactions", controller.Authenticate(controller.InsertNewTransactions, 1)).Methods("POST")

	//products
	router.HandleFunc("/products", controller.GetAllProduct).Methods("GET")
	router.HandleFunc("/products", controller.Authenticate(controller.InsertNewProduct, 1)).Methods("POST")
	router.HandleFunc("/products/update/{id}", controller.Authenticate(controller.UpdateProduct, 1)).Methods("PUT")
	router.HandleFunc("/products/delete/{id}", controller.Authenticate(controller.DeleteProduct, 2)).Methods("DELETE")
	//feedback
	router.HandleFunc("/feedbacks", controller.Authenticate(controller.GetFeedback, 1)).Methods("GET")
	router.HandleFunc("/feedbacks", controller.Authenticate(controller.InsertNewFeedback, 1)).Methods("POST")
	router.HandleFunc("/feedbacks/all", controller.Authenticate(controller.GetAllFeedbacks, 3)).Methods("GET")

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
