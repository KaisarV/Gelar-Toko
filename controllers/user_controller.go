package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	config "GelarToko/config"
	model "GelarToko/models"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()

	defer db.Close()
	var response model.ErrorResponse

	err := r.ParseForm()
	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		log.Println(err.Error())
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if email == "" || password == "" {
		response.Status = 400
		response.Message = "Please input name and password"

		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	rows, err := db.Query("SELECT * FROM users WHERE email=? AND password=?",
		email,
		password,
	)

	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	var users []model.User

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Email, &user.Password, &user.Address, &user.UserType, &user.IsVerified); err != nil {
			log.Print(err.Error())
		} else {
			users = append(users, user)
		}
	}

	if len(users) == 1 {
		generateToken(w, user.ID, user.Name, user.UserType)

		response.Status = 200
		response.Message = "Login Success"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {

		response.Status = 400
		response.Message = "Login Failed"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)

	var response model.UserResponse
	response.Status = 200
	response.Message = "Logout Success"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
