package controllers

import (
	config "GelarToko/config"
	model "GelarToko/models"
	"fmt"
	"net/http"
	"strconv"
)

func InsertNewTransactions(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var response model.TransactionResponse
	var transaction model.Transaction
	_, transaction.UserId, _, _ = validateTokenFromCookies(r)

	err := r.ParseForm()
	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		SendResponse(w, response.Status, response)
		return
	}

	transaction.ProductId, _ = strconv.Atoi(r.Form.Get("ProductID"))
	transaction.Quantity, _ = strconv.Atoi(r.Form.Get("Quantity"))

	rows, _ := db.Query("SELECT Stock FROM products WHERE Id = ?", transaction.ProductId)

	var product model.Product
	var products []model.Product

	i := 0
	for rows.Next() {
		if err := rows.Scan(&product.Stock); err != nil {
		} else {
			products = append(products, product)
		}
		i++
	}

	if i == 0 {
		response.Status = 400
		response.Message = "Product not found"
		SendResponse(w, response.Status, response)
		return
	}

	if products[0].Stock < transaction.Quantity {
		response.Status = 400
		response.Message = "Insufficient product stock"
		SendResponse(w, response.Status, response)
		return
	}

	fmt.Println(products[0].Stock)

	res, errQuery := db.Exec("INSERT INTO transactions (User_Id, Product_Id, Quantity) VALUES (?, ?, ?)", transaction.UserId, transaction.ProductId, transaction.Quantity)
	// print(errQuery.Error())
	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		transaction.ID = int(id)
		response.Data = transaction
		SendResponse(w, response.Status, response)
		newStock := product.Stock - transaction.Quantity
		_, err := db.Exec("UPDATE Products SET Stock = ? WHERE Id = ? ", newStock, transaction.ProductId)

		if err != nil {
			response.Status = 400
			response.Message = "Error Update Product"
			SendResponse(w, response.Status, response)
			return
		}

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
	fmt.Print(userId)
	err := r.ParseForm()
	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		SendResponse(w, response.Status, response)
		return
	}
	// convertedString := strconv.Itoa(userId)

	// query := "SELECT * FROM  transactions  WHERE User_Id  = " + convertedString
	rows, err := db.Query("SELECT * FROM  transactions  WHERE User_Id =? ", userId)

	if err != nil {
		response.Status = 400
		response.Message = "Error Query"
		SendResponse(w, response.Status, response)
		return
	}

	for rows.Next() {
		if err := rows.Scan(&transaction.ID, &transaction.UserId, &transaction.ProductId, &transaction.Date, &transaction.Price, &transaction.Quantity, &transaction.Status); err != nil {
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

func UpdateTransactions(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var response model.TransactionsResponse
	// var transaction model.Transaction
	var transactions []model.Transaction

	_, userId, _, _ := validateTokenFromCookies(r)
	fmt.Print(userId)
	err := r.ParseForm()
	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		SendResponse(w, response.Status, response)
	}

	transactionsID := r.Form.Get("transID")
	Status := r.Form.Get("status")
	//0 -> packed
	//1 -> sent
	//2 -> arrived
	// Status, _ := strconv.Atoi(r.Form.Get("status"))

	_, errQuery := db.Exec("UPDATE transactions set Status ='" + Status + "' WHERE Id = " + transactionsID)

	if errQuery == nil {
		response.Status = 200
		response.Message = "Succes Get Data"
		response.Data = transactions
		SendResponse(w, response.Status, response)
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
		SendResponse(w, response.Status, response)
	}
}
