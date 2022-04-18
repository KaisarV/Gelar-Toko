package controllers

import (
	config "GelarToko/config"
	model "GelarToko/models"
	"fmt"
	"log"
	"net/http"
)

func GetAllStores(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	var response model.StoresResponse
	defer db.Close()

	query := "SELECT * FROM stores"
	id := r.URL.Query()["id"]
	if id != nil {
		query += " WHERE id = " + id[0]
	}

	rows, err := db.Query(query)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		SendResponse(w, response.Status, response)
		return
	}

	var store model.Store
	var stores []model.Store

	for rows.Next() {
		if err := rows.Scan(&store.ID, &store.UserId, &store.Name, &store.Address); err != nil {
			log.Println(err.Error())
		} else {
			stores = append(stores, store)
		}
	}

	if len(stores) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = stores
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}

	SendResponse(w, response.Status, response)
}

func DeleteMyStore(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var response model.ErrorResponse
	_, userId, userName, _ := validateTokenFromCookies(r)
	fmt.Println("Disini1")
	query, errQuery := db.Exec(`DELETE FROM stores WHERE User_Id = ?;`, userId)
	RowsAffected, _ := query.RowsAffected()
	fmt.Println("Disini2")
	if RowsAffected == 0 {

		response.Status = 400
		response.Message = "store not found"
		SendResponse(w, response.Status, response)
		return
	}
	fmt.Println("Disini3")
	if errQuery == nil {
		query, _ = db.Exec("UPDATE users SET User_Type = ? WHERE Id = ?", 1, userId)
		response.Status = 200
		response.Message = "Success Delete Data"
		generateToken(w, userId, userName, 1)
	} else {
		response.Status = 400
		response.Message = "Error Delete Data"
		log.Println(errQuery.Error())
	}
	SendResponse(w, response.Status, response)
}

func InsertMyStore(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()

	var store model.Store
	var response model.StoreResponse
	var userType int
	var userName string

	_, store.UserId, userName, userType = validateTokenFromCookies(r)

	if userType == 2 {
		response.Status = 400
		response.Message = "You already have a store"
		SendResponse(w, response.Status, response)
		return
	}

	err := r.ParseForm()

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		log.Println(err.Error())
		return
	}

	store.Name = r.Form.Get("name")
	store.Address = r.Form.Get("address")

	// store.Name = input.Name
	// store.Address = input.Address

	if store.Name == "" {
		response.Status = 400
		response.Message = "Please Insert store's Name"
		SendResponse(w, response.Status, response)
		return
	}

	if store.Address == "" {
		response.Status = 400
		response.Message = "Please Insert store's Address"
		SendResponse(w, response.Status, response)
		return
	}

	res, errQuery := db.Exec("INSERT INTO stores(User_Id, Name,  Address) VALUES(?, ?, ?)", store.UserId, store.Name, store.Address)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		res, _ = db.Exec("UPDATE users SET User_Type = ? WHERE Id = ?", 2, store.UserId)
		response.Status = 200
		response.Message = "Success"
		store.ID = int(id)
		response.Data = store
		generateToken(w, store.UserId, userName, 2)
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}
	SendResponse(w, response.Status, response)
}

func UpdateMyStore(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()

	var store model.Store
	var response model.StoreResponse
	err := r.ParseForm()

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		log.Println(err.Error())
		return
	}

	store.Name = r.Form.Get("name")
	store.Address = r.Form.Get("address")

	_, userId, _, _ := validateTokenFromCookies(r)
	// store.Name = input.Name
	// store.Address = input.Address

	rows, _ := db.Query("SELECT * FROM stores WHERE User_Id = ?", userId)
	var prevDatas []model.Store
	var prevData model.Store

	for rows.Next() {
		if err := rows.Scan(&prevData.ID, &prevData.UserId, &prevData.Name, &prevData.Address); err != nil {
			log.Println(err.Error())
		} else {
			prevDatas = append(prevDatas, prevData)
		}
	}

	if len(prevDatas) > 0 {
		if store.Name == "" {
			store.Name = prevDatas[0].Name
		}
		if store.Address == "" {
			store.Address = prevDatas[0].Address
		}

		_, errQuery := db.Exec(`UPDATE stores SET Name = ?,  Address = ? WHERE User_id = ?`, store.Name, store.Address, userId)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Success Update Data"
			store.ID = prevData.ID
			response.Data = store
		} else {
			response.Status = 400
			response.Message = "Error Update Data"

			log.Println(errQuery)
		}
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	SendResponse(w, response.Status, response)
}
