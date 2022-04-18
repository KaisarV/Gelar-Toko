package controllers

import (
	model "GelarToko/models"
	"encoding/json"
	"net/http"
)

func SendUnAuthorizedResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	var response model.MessageResponse
	response.Status = status
	response.Message = message
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func SendResponse(w http.ResponseWriter, status int, response interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
