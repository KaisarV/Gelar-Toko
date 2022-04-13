package controllers

import (
	config "GelarToko/config"
	model "GelarToko/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SendChat(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var response model.ChatResponse
	err := r.ParseForm()

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		log.Println(err.Error())
		return
	}

	var chat model.Chat
	vars := mux.Vars(r)

	_, _, _, userTypeFromCookie := validateTokenFromCookies(r)
	var userType int
	if userTypeFromCookie == 1 {
		userType = 1
	} else {
		userType, _ = strconv.Atoi(vars["userType"])
	}
	chat.ReceiverId, _ = strconv.Atoi(vars["receiverId"])
	var customerId int

	if userType == 1 {
		_, customerId, _, _ = validateTokenFromCookies(r)

		rows, _ := db.Query("SELECT User_Type FROM users WHERE Id = ?", customerId)

		var users []model.User
		var user model.User
		for rows.Next() {
			if err := rows.Scan(&user.UserType); err != nil {
				log.Println(err.Error())
			} else {
				users = append(users, user)
			}
		}

		if users[0].UserType == 1 {
			response.Status = 200
			response.Message = "User doesn't have store"
			SendResponse(w, response.Status, response)
		}

	} else if userType == 2 {
		customerId = chat.ReceiverId
	} else {
		response.Status = 200
		response.Message = "Error UserType Parameter"
		SendResponse(w, response.Status, response)
	}

	_, chat.SenderId, _, _ = validateTokenFromCookies(r)
	chat.Chat = r.Form.Get("chat")

	res, errQuery := db.Exec("INSERT INTO chat(Sender_Id, Receiver_Id, Chat, Customer_Id) VALUES(?, ?, ?, ?)", chat.SenderId, chat.ReceiverId, chat.Chat, customerId)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		chat.ID = int(id)
		response.Data = chat
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}

	SendResponse(w, response.Status, response)
}

func GetChat(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	_, _, _, userTypeFromCookie := validateTokenFromCookies(r)
	var userType int
	vars := mux.Vars(r)

	if userTypeFromCookie == 1 {
		userType = 1
	} else {
		userType, _ = strconv.Atoi(vars["userType"])
	}

	var response model.ChatsResponse

	if userType == 1 {
		_, customerId, _, _ := validateTokenFromCookies(r)

		rows, _ := db.Query("SELECT * FROM chat WHERE Customer_Id = ?", customerId)

		var chats []model.Chat
		var chat model.Chat
		for rows.Next() {
			if err := rows.Scan(&chat.ID, &chat.SenderId, &chat.ReceiverId, &chat.CustomerId, &chat.Chat, &chat.Date); err != nil {
				log.Println(err.Error())
			} else {
				chats = append(chats, chat)
			}
		}

		if len(chats) != 0 {
			response.Status = 200
			response.Message = "Success Get Data"
			response.Data = chats
		} else {
			response.Status = 400
			response.Message = "Error Get Data"
			w.WriteHeader(400)
		}

	} else if userType == 2 {
		_, userId, _, _ := validateTokenFromCookies(r)
		rows, _ := db.Query("SELECT * FROM chat WHERE Customer_Id != ? AND Sender_Id = ? OR Receiver_Id = ?", userId, userId, userId)

		var chats []model.Chat
		var chat model.Chat
		for rows.Next() {
			if err := rows.Scan(&chat.ID, &chat.SenderId, &chat.ReceiverId, &chat.CustomerId, chat.Chat, chat.Date); err != nil {
				log.Println(err.Error())
			} else {
				chats = append(chats, chat)
			}
		}

		if len(chats) != 0 {
			response.Status = 200
			response.Message = "Success Get Data"
			response.Data = chats
		} else {
			response.Status = 400
			response.Message = "Error Get Data"
			w.WriteHeader(400)
		}
	} else {
		response.Status = 200
		response.Message = "Error UserType Parameter"
		SendResponse(w, response.Status, response)
	}

	SendResponse(w, response.Status, response)
}
