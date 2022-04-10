package controllers

import (
	config "GelarToko/config"
	model "GelarToko/models"
	"net/http"
	"strconv"
)

func InsertNewTransactions(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var response model.TransactionResponse
	var transaction model.Transaction
	_, userId, _, _ := validateTokenFromCookies(r)

	transaction.ProductId, _ = strconv.Atoi(r.Form.Get("ProductID"))
	transaction.Quantity, _ = strconv.Atoi(r.Form.Get("Quantity"))

	err := r.ParseForm()
	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		SendResponse(w, response.Status, response)
		return
	}

	_, errQuery := db.Exec("INSERT INTO transactions (userid, productid, quantity) VALUES (?,?,?)", userId, transaction.ProductId, transaction.Quantity)

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		response.Data = transaction
		SendResponse(w, response.Status, response)
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		SendResponse(w, response.Status, response)
	}

}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var response model.TransactionsResponse
	var transaction model.Transaction
	var transactions []model.Transaction
	_, userId, _, _ := validateTokenFromCookies(r)

	err := r.ParseForm()
	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		SendResponse(w, response.Status, response)
		return
	}
	convertedString := strconv.Itoa(userId)

	query := "SELECT * FROM  transactions  WHERE User_id  = " + convertedString
	rows, err := db.Query(query)

	if err != nil {
		response.Status = 400
		response.Message = "Error Query"
		SendResponse(w, response.Status, response)
		return
	}

	for rows.Next() {
		if err := rows.Scan(&transaction.ID, &transaction.UserId, &transaction.ProductId, &transaction.Quantity); err != nil {
		} else {
			transactions = append(transactions, transaction)
		}
	}

	if len(transactions) != 0 {
		response.Status = 200
		response.Message = "Succes Get Data"
		response.Data = transactions
		SendResponse(w, response.Status, response)
	} else if response.Message == "" {
		response.Status = 400
		response.Message = "Data Not Found"
		SendResponse(w, response.Status, response)
	}

}
