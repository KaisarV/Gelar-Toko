package controllers

import (
	config "GelarToko/config"
	model "GelarToko/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllProduct(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	var response model.ProductsResponse
	defer db.Close()

	query := "SELECT Id, Name, Category, Price FROM products"
	name := r.URL.Query()["name"]
	if name != nil {
		query += " WHERE Name LIKE % " + name[0] + "%"
	}

	rows, err := db.Query(query)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		SendResponse(w, response.Status, response)
		return
	}

	var product model.Product
	var products []model.Product

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price); err != nil {

		} else {
			products = append(products, product)
		}
	}

	if len(products) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = products
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	SendResponse(w, response.Status, response)
}

func InsertNewProduct(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var store model.Store
	var stores []model.Store
	var product model.Product
	var response model.ProductResponse

	_, store.UserId, _, _ = validateTokenFromCookies(r)

	err := r.ParseForm()

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		return
	}

	fmt.Println(store.UserId)
	//check store id by userID
	rows, _ := db.Query("SELECT Id FROM stores WHERE User_Id = ?", store.UserId)
	for rows.Next() {
		if err := rows.Scan(&store.ID); err != nil {
			response.Status = 400
			response.Message = "Internal Error"
			SendResponse(w, response.Status, response)
		} else {
			stores = append(stores, store)
		}

	}

	product.Name = r.Form.Get("name")
	product.Category = r.Form.Get("Category")
	product.Price, _ = strconv.Atoi(r.Form.Get("Price"))

	fmt.Println(stores[0].ID)

	if product.Name == "" {
		response.Status = 400
		response.Message = "Please Insert Product Name"
		SendResponse(w, response.Status, response)
		return
	}

	if product.Category == "" {
		response.Status = 400
		response.Message = "Please Insert Category"
		SendResponse(w, response.Status, response)
		return
	}

	if product.Price == 0 {
		response.Status = 400
		response.Message = "Please Insert Price "
		SendResponse(w, response.Status, response)
		return
	}

	_, errQuery := db.Exec("INSERT INTO products (Name, Category, Price , Store_Id) VALUES (?,?,?,?)", product.Name, product.Category, product.Price, stores[0].ID)
	// fmt.Print(errQuery.Error())
	if errQuery != nil {
		response.Status = 400
		response.Message = "Error Insert Data"
		SendResponse(w, response.Status, response)
	} else {
		response.Status = 200
		response.Message = "Success"
		response.Data = product
		SendResponse(w, response.Status, response)
	}

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var store model.Store
	var stores []model.Store
	var product model.Product
	var response model.ProductResponse

	_, store.UserId, _, _ = validateTokenFromCookies(r)

	err := r.ParseForm()

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		return
	}

	fmt.Println(store.UserId)
	//check store id by userID
	rows, _ := db.Query("SELECT Id FROM stores WHERE User_Id = ?", store.UserId)
	for rows.Next() {
		if err := rows.Scan(&store.ID); err != nil {
			response.Status = 400
			response.Message = "Internal Error"
			SendResponse(w, response.Status, response)
		} else {
			stores = append(stores, store)
		}

	}

	vars := mux.Vars(r)
	product.ID, _ = strconv.Atoi(vars["id"])
	product.Name = r.Form.Get("name")
	product.Category = r.Form.Get("Category")
	product.Price, _ = strconv.Atoi(r.Form.Get("Price"))

	fmt.Println(stores[0].ID)

	if product.Name == "" {
		response.Status = 400
		response.Message = "Please Insert Product Name"
		SendResponse(w, response.Status, response)
		return
	}

	if product.Category == "" {
		response.Status = 400
		response.Message = "Please Insert Category"
		SendResponse(w, response.Status, response)
		return
	}

	if product.Price == 0 {
		response.Status = 400
		response.Message = "Please Insert Price "
		SendResponse(w, response.Status, response)
		return
	}

	_, errQuery := db.Exec(`UPDATE products SET Name = ?,  Category = ?, Price = ?, Store_Id = ? WHERE id = ?`, product.Name, product.Category, product.Price, stores[0].ID, product.ID)
	// fmt.Print(errQuery.Error())
	if errQuery != nil {
		response.Status = 400
		response.Message = "Error Insert Data"
		SendResponse(w, response.Status, response)
	} else {
		response.Status = 200
		response.Message = "Success"
		response.Data = product
		SendResponse(w, response.Status, response)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	defer db.Close()

	var store model.Store
	var stores []model.Store
	var product model.Product
	var response model.ProductResponse

	_, store.UserId, _, _ = validateTokenFromCookies(r)

	err := r.ParseForm()

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		SendResponse(w, response.Status, response)
	}

	//cek id toko
	rows, _ := db.Query("SELECT Id FROM stores WHERE User_Id = ?", store.UserId)
	for rows.Next() {
		if err := rows.Scan(&store.ID); err != nil {
			response.Status = 400
			response.Message = "Internal Error"
			SendResponse(w, response.Status, response)
		} else {
			stores = append(stores, store)
		}

	}

	vars := mux.Vars(r)
	product.ID, _ = strconv.Atoi(vars["id"])
	//delete barang yang ada di toko tersebut
	query, errQuery := db.Exec(`DELETE FROM products WHERE id = ? AND Store_id = ?;`, product.ID, stores[0].ID)
	RowsAffected, _ := query.RowsAffected()

	if RowsAffected == 0 {
		response.Status = 400
		response.Message = "There is no such product on the store"
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
	}
	SendResponse(w, response.Status, response)
}

func BlockirProduct(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	defer db.Close()

}
