package controllers

import (
	model "GelarToko/models"
	"encoding/json"
	"net/http"
)

func SendUnAuthorizedResponse(w http.ResponseWriter, message string, status int) {
	var response model.MessageResponse
	response.Status = status
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
