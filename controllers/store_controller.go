package controllers

import (
	config "GelarToko/config"
	model "GelarToko/models"
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
	return
}
