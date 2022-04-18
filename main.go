package main

import (
	controller "GelarToko/controllers"
	"fmt"
	"log"
	"net/http"
	"time"

	config "GelarToko/config"
	model "GelarToko/models"

	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// for load godotenv
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	//1. User Biasa 2. Memiliki Toko 3. Admin
	router.HandleFunc("/login", controller.UserLogin).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")
	router.HandleFunc("/register", controller.InsertUser).Methods("POST")
	//Verifikasi akun setelah login dengan link yang dikirim ke email
	router.HandleFunc("/verify/{token}", controller.VerifyToken).Methods("GET")
	//Verifikasi akun manual(untuk testing tanpa lewat email(email dummy))
	router.HandleFunc("/verify/testing/{id}", controller.VerifyTokenById).Methods("GET")

	//User
	router.HandleFunc("/users", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controller.Authenticate(controller.DeleteUser, 3)).Methods("DELETE")
	router.HandleFunc("/users", controller.Authenticate(controller.UpdateMyProfile, 1)).Methods("PUT")
	router.HandleFunc("/user/block/{id}", controller.Authenticate(controller.BlockUser, 3)).Methods("GET")
	router.HandleFunc("/user/unblock/{id}", controller.Authenticate(controller.UnblockUser, 3)).Methods("GET")

	//Store
	router.HandleFunc("/stores", controller.GetAllStores).Methods("GET")
	router.HandleFunc("/stores", controller.Authenticate(controller.InsertMyStore, 1)).Methods("POST")
	router.HandleFunc("/stores", controller.Authenticate(controller.DeleteMyStore, 2)).Methods("DELETE")
	router.HandleFunc("/stores", controller.Authenticate(controller.UpdateMyStore, 2)).Methods("PUT")

	//Chat
	router.HandleFunc("/chat/{receiverId}/{userType}", controller.Authenticate(controller.SendChat, 1)).Methods("POST")
	router.HandleFunc("/chat/{userType}", controller.Authenticate(controller.GetChat, 1)).Methods("GET")

	//transactions
	router.HandleFunc("/transactions", controller.Authenticate(controller.GetTransaction, 1)).Methods("GET")
	router.HandleFunc("/transaction", controller.Authenticate(controller.InsertNewTransactions, 1)).Methods("POST")
	router.HandleFunc("/transaction", controller.Authenticate(controller.UpdateTransactions, 2)).Methods("PUT")

	//products
	router.HandleFunc("/products", controller.GetAllProduct).Methods("GET")
	router.HandleFunc("/products", controller.Authenticate(controller.InsertNewProduct, 1)).Methods("POST")
	router.HandleFunc("/products/update/{id}", controller.Authenticate(controller.UpdateProduct, 1)).Methods("PUT")
	router.HandleFunc("/products/delete/{id}", controller.Authenticate(controller.DeleteProduct, 2)).Methods("DELETE")
	router.HandleFunc("/products/block/{id}", controller.Authenticate(controller.BlockProduct, 3)).Methods("PUT")
	router.HandleFunc("/products/unblock/{id}", controller.Authenticate(controller.UnblockProduct, 3)).Methods("PUT")
	router.HandleFunc("/products/get-blocked-store-products", controller.Authenticate(controller.GetBlockedStoreProduct, 3)).Methods("GET")

	//feedback
	router.HandleFunc("/feedbacks", controller.Authenticate(controller.GetFeedback, 1)).Methods("GET")
	router.HandleFunc("/feedbacks", controller.Authenticate(controller.InsertNewFeedback, 1)).Methods("POST")
	router.HandleFunc("/feedbacks/all", controller.Authenticate(controller.GetAllFeedbacks, 3)).Methods("GET")

	//Product Review
	router.HandleFunc("/review/{productid}", controller.Authenticate(controller.InsertMyProductReview, 1)).Methods("POST")
	router.HandleFunc("/reviews", controller.Authenticate(controller.GetAllMyProductReviews, 1)).Methods("GET")
	router.HandleFunc("/review/{productid}", controller.Authenticate(controller.DeleteMyProductReview, 1)).Methods("DELETE")
	router.HandleFunc("/review/{productid}", controller.Authenticate(controller.UpdateMyProductReview, 1)).Methods("PUT")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	s := gocron.NewScheduler(time.UTC)

	s.Every(1).MonthLastDay().Do(func() {
		db := config.Connect()
		defer db.Close()

		var response model.ErrorResponse
		_, err := db.Exec("UPDATE products SET Current_Price = Normal_Price * 0.8")

		if err != nil {
			response.Status = 400
			response.Message = err.Error()

		} else {
			response.Status = 200
			response.Message = "Success update price"
		}
		fmt.Println(response)
	})

	s.Every(1).Months(1).Do(func() {
		db := config.Connect()
		defer db.Close()

		var response model.ErrorResponse
		_, err := db.Exec("UPDATE products SET Current_Price = Normal_Price")

		if err != nil {
			response.Status = 400
			response.Message = err.Error()

		} else {
			response.Status = 200
			response.Message = "Success update price"
		}
		fmt.Println(response)
	})

	s.StartAsync()
	handler := corsHandler.Handler(router)
	log.Println("Starting on Port")

	err := http.ListenAndServe(":8080", handler)
	log.Fatal(err)
}
