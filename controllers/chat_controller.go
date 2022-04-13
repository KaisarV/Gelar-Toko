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
	userType, _ := strconv.Atoi(vars["userType"])
	chat.ReceiverId, _ = strconv.Atoi(vars["receiverId"])
	var customerId int

	if userType == 1 {
		_, customerId, _, _ = validateTokenFromCookies(r)
	} else if userType == 2 {
		customerId = chat.ReceiverId
	} else {
		response.Status = 200
		response.Message = "Error UserType Parameter"
		SendResponse(w, response.Status, response)
	}

	_, chat.SenderId, _, _ = validateTokenFromCookies(r)
	chat.Chat = r.Form.Get("chat")

	res, errQuery := db.Exec("INSERT INTO chat(Sender_Id, Receiver_Id, Chat, Customer_Id) VALUES(?, ?, ?)", chat.SenderId, chat.ReceiverId, customerId)

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

// func GetChat(w http.ResponseWriter, r *http.Request) {
// 	db := config.Connect()
// 	defer db.Close()

// }
