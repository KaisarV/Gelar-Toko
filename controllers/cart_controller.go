package controllers

import (
	config "GelarToko/config"
	model "GelarToko/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCartItem(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var response model.CartsResponse
	var cart model.Cart
	var carts []model.Cart

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
	rows, err := db.Query(`SELECT * FROM  carts  WHERE User_Id = ? `, userId)

	if err != nil {
		response.Status = 400
		response.Message = "Error Query"
		SendResponse(w, response.Status, response)
		return
	}

	for rows.Next() {
		if err := rows.Scan(&cart.ID, &cart.UserId, &cart.ProductId, &cart.Quantity); err != nil {
		} else {
			carts = append(carts, cart)
		}
	}

	if len(carts) != 0 {
		response.Status = 200
		response.Message = "Succes Get Data"
		response.Data = carts
		SendResponse(w, response.Status, response)
	} else if response.Message == "" {
		response.Status = 400
		response.Message = "Data Not Found"
		SendResponse(w, response.Status, response)
	}

}

func InsertCartItem(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var response model.CartsResponse
	var cart model.Cart
	var carts []model.Cart
	_, cart.UserId, _, _ = validateTokenFromCookies(r)

	err := r.ParseForm()
	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		SendResponse(w, response.Status, response)
		return
	}

	// cart.UserId, _ = strconv.Atoi(r.Form.Get("UserID"))
	cart.ProductId, _ = strconv.Atoi(r.Form.Get("ProductID"))
	cart.Quantity, _ = strconv.Atoi(r.Form.Get("Quantity"))

	res, errQuery := db.Exec(`INSERT INTO carts (User_Id, Product_Id, Quantity) VALUES (?, ?, ?)`,
		cart.UserId,
		cart.ProductId,
		cart.Quantity,
	)
	// print(errQuery.Error())
	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		cart.ID = int(id)
		response.Data = carts // klo gk pake yg array jadi error gk tau knapa ni
		SendResponse(w, response.Status, response)
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		SendResponse(w, response.Status, response)
	}

}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	err := r.ParseForm()
	var response model.ErrorResponse
	var cart model.Cart

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		log.Println(err.Error())
		return
	}

	vars := mux.Vars(r)
	_, cart.UserId, _, _ = validateTokenFromCookies(r) // 2 param buat delet
	fmt.Print(cart.UserId)
	cartId := vars["cartId"] //userid -> kerangjang siapa cartid -> cart yang mana
	query, errQuery := db.Exec(`DELETE FROM carts WHERE User_Id = ? And Id = ?`,
		cart.UserId,
		cartId,
	)
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

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	err := r.ParseForm()
	var response model.CartsResponse
	var carts []model.Cart
	var cart model.Cart

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		log.Println(err.Error())
		return
	}

	_, cart.UserId, _, _ = validateTokenFromCookies(r)
	cart.ID, _ = strconv.Atoi(r.Form.Get("Id"))
	// cart.ProductId, _ = strconv.Atoi(r.Form.Get("ProdId"))
	cart.Quantity, _ = strconv.Atoi(r.Form.Get("qty"))

	_, errQuery := db.Exec(
		`UPDATE carts SET Quantity = ? WHERE User_Id = ? AND Id = ?`, //AND Product_Id = ?
		cart.Quantity,
		cart.UserId,
		/*cart.ProductId,*/
		cart.ID,
	)

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		response.Data = carts
		SendResponse(w, response.Status, response)
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
		SendResponse(w, response.Status, response)
	}
}
