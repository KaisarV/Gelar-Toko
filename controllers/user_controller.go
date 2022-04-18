package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	config "GelarToko/config"
	gomail "GelarToko/gomail"
	model "GelarToko/models"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	var response model.UsersResponse
	defer db.Close()

	query := "SELECT Id, Name, Phone, Email,  Address, User_Type FROM users WHERE is_Verified != -1"
	id := r.URL.Query()["id"]
	if id != nil {
		query += " AND id = " + id[0]
	}

	rows, err := db.Query(query)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		SendResponse(w, response.Status, response)
		return
	}

	var user model.User
	var users []model.User

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Email, &user.Address, &user.UserType); err != nil {
			log.Println(err.Error())
		} else {
			users = append(users, user)
		}
	}

	if len(users) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = users
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	SendResponse(w, response.Status, response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	err := r.ParseForm()
	var response model.ErrorResponse

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		log.Println(err.Error())
		return
	}

	vars := mux.Vars(r)
	userId := vars["id"]
	query, errQuery := db.Exec(`DELETE FROM users WHERE id = ?;`, userId)
	RowsAffected, _ := query.RowsAffected()

	if RowsAffected == 0 {
		response.Status = 400
		response.Message = "User not found"
		SendResponse(w, response.Status, response)
		return
	}

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success Delete Data"
	} else {
		response.Status = 400
		response.Message = "Error Delete Data"
		w.WriteHeader(400)
		log.Println(errQuery.Error())
	}
	SendResponse(w, response.Status, response)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {

	_, userId, _, _ := validateTokenFromCookies(r)
	var response model.UserResponse

	if userId != -1 {
		response.Status = 400
		response.Message = "You are already logged in"
		w.WriteHeader(400)
		SendResponse(w, response.Status, response)
		return
	}

	db := config.Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		log.Println(err.Error())
		return
	}

	var user model.User

	user.Name = r.Form.Get("name")
	user.Phone = r.Form.Get("phone")
	user.Email = r.Form.Get("email")
	user.Password = r.Form.Get("password")
	user.Address = r.Form.Get("address")

	if user.Name == "" {
		response.Status = 400
		response.Message = "Please Insert User's Name"
		SendResponse(w, response.Status, response)
		return
	}

	if user.Phone == "" {
		response.Status = 400
		response.Message = "Please Insert User's Phone"
		SendResponse(w, response.Status, response)
		return
	}

	if user.Address == "" {
		response.Status = 400
		response.Message = "Please Insert User's Address"
		SendResponse(w, response.Status, response)
		return
	}

	if user.Email == "" {
		response.Status = 400
		response.Message = "Please Insert User's Email"
		SendResponse(w, response.Status, response)
		return
	}

	if user.Password == "" {
		response.Status = 400
		response.Message = "Please Insert User's Password"
		SendResponse(w, response.Status, response)
		return
	}

	// user.Name = input.Name
	// user.Phone = input.Phone
	// user.Email = input.Email
	// user.Password = input.Password
	// user.Address = input.Address

	rows, _ := db.Query("SELECT Email FROM users WHERE Email = ?", user.Email)

	i := 0
	for rows.Next() {
		i++
	}

	if i != 0 {
		response.Status = 400
		response.Message = "Email already registered"
		SendResponse(w, response.Status, response)
		return
	}

	res, errQuery := db.Exec("INSERT INTO users(Name, Phone,  Email, Password,Address) VALUES(?, ?, ?, ?, ?)", user.Name, user.Phone, user.Email, user.Password, user.Address)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		user.UserType = 1
		user.ID = int(id)
		response.Data = user
		gomail.SendRegisterMail(user.Email, user.Name, user.ID)
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}

	SendResponse(w, response.Status, response)
}

func UpdateMyProfile(w http.ResponseWriter, r *http.Request) {

	var user model.User
	var response model.UserResponse

	db := config.Connect()
	defer db.Close()
	err := r.ParseForm()

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		log.Println(err.Error())
		return
	}

	user.Name = r.Form.Get("name")
	user.Phone = r.Form.Get("phone")
	user.Email = r.Form.Get("email")
	user.Password = r.Form.Get("password")
	user.Address = r.Form.Get("address")

	_, userId, _, _ := validateTokenFromCookies(r)
	log.Println()
	// user.Name = input.Name
	// user.Phone = input.Phone
	// user.Email = input.Email
	// user.Password = input.Password
	// user.Address = input.Address

	rows, _ := db.Query("SELECT * FROM users WHERE Id = ?", userId)
	var prevDatas []model.User
	var prevData model.User

	for rows.Next() {
		if err := rows.Scan(&prevData.ID, &prevData.Name, &prevData.Phone, &prevData.Email, &prevData.Password, &prevData.Address, &prevData.UserType, &prevData.IsVerified); err != nil {
			log.Println(err.Error())
		} else {
			prevDatas = append(prevDatas, prevData)
		}
	}

	if len(prevDatas) > 0 {
		if user.Name == "" {
			user.Name = prevDatas[0].Name
		}
		if user.Address == "" {
			user.Address = prevDatas[0].Address
		}
		if user.Phone == "" {
			user.Phone = prevDatas[0].Phone
		}
		if user.Email == "" {
			user.Email = prevDatas[0].Email
		}
		if user.Password == "" {
			user.Password = prevDatas[0].Password
		}

		_, errQuery := db.Exec(`UPDATE users SET Name = ?,  Phone = ?, Email = ?, Password = ?,Address = ? WHERE id = ?`, user.Name, user.Phone, user.Email, user.Password, user.Address, userId)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Success Update Data"
			user.ID = userId
			response.Data = user
		} else {
			response.Status = 400
			response.Message = "Error Update Data"
			w.WriteHeader(400)
			log.Println(errQuery)
		}
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
		w.WriteHeader(400)
	}
	SendResponse(w, response.Status, response)

	if user.Name != "" {
		generateToken(w, user.ID, user.Name, user.UserType)
	}
}

func BlockUser(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	userId := vars["id"]

	var response model.ErrorResponse
	var user model.User

	_, errQuery := db.Exec(`UPDATE users SET Is_Verified = ? WHERE id = ?`, -1, userId)

	if errQuery == nil {
		generateToken(w, user.ID, user.Name, user.UserType)
		response.Status = 200
		response.Message = "Success Block User"
		go gomail.SendBlockMail(user.Email, user.Name)

	} else {
		response.Status = 400
		response.Message = "Error Block User"
	}
	SendResponse(w, response.Status, response)
}

func UnblockUser(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	userId := vars["id"]

	var response model.ErrorResponse
	var user model.User

	_, errQuery := db.Exec(`UPDATE users SET Is_Verified = ? WHERE id = ?`, 1, userId)

	if errQuery == nil {
		generateToken(w, user.ID, user.Name, user.UserType)
		response.Status = 200
		response.Message = "Success Unblock User"
		go gomail.SendUnblockMail(user.Email, user.Name)

	} else {
		response.Status = 400
		response.Message = "Error Unblock User"
	}
	SendResponse(w, response.Status, response)
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	token := vars["token"]
	userId := GetDataFromToken(token)
	var response model.ErrorResponse

	if userId == -1 {
		response.Status = 400
		response.Message = "Id Not Found"
		SendResponse(w, response.Status, response)
		return
	}

	db := config.Connect()
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM users WHERE Id = ?", userId)

	i := 0
	for rows.Next() {
		i++
	}

	if i == 0 {
		response.Status = 400
		response.Message = "Account Not Found"
		SendResponse(w, response.Status, response)
		return
	}

	_, errQuery := db.Exec(`UPDATE users SET Is_Verified = ? WHERE id = ?`, 1, userId)

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success Verify Account"
	} else {
		response.Status = 400
		response.Message = "Failed Verify Account"
		log.Println(errQuery)
	}

	SendResponse(w, response.Status, response)
}

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

	rows, err := db.Query("SELECT * FROM users WHERE email=? AND password=? AND Is_Verified != -1",
		email,
		password,
	)

	if err != nil {
		response.Status = 400
		response.Message = "Data Found"
		SendResponse(w, response.Status, response)
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
		go gomail.SendLoginMail(user.Email, user.Name)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		response.Status = 400
		response.Message = "Login Failed"
		w.WriteHeader(400)
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

func VerifyTokenById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["id"]

	var response model.ErrorResponse

	db := config.Connect()
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM users WHERE Id = ?", userId)

	i := 0
	for rows.Next() {
		i++
	}

	if i == 0 {
		response.Status = 400
		response.Message = "Account Not Found"
		SendResponse(w, response.Status, response)
		return
	}

	_, errQuery := db.Exec(`UPDATE users SET Is_Verified = ? WHERE id = ?`, 1, userId)

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success Verify Account"
	} else {
		response.Status = 400
		response.Message = "Failed Verify Account"
		log.Println(errQuery)
	}

	SendResponse(w, response.Status, response)
}
