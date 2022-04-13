package controllers

import (
	config "GelarToko/config"
	model "GelarToko/models"
	"fmt"
	"log"
	"net/http"
)

func InsertNewFeedback(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var response model.FeedbackResponse
	var feedback model.Feedback
	_, userId, _, _ := validateTokenFromCookies(r)

	err := r.ParseForm()
	feedback.Feedback = r.Form.Get("Feedback")
	fmt.Print(feedback.Feedback)

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		SendResponse(w, response.Status, response)
		return
	}

	_, errQuery := db.Exec("INSERT INTO feedbacks (user_id, feedback) VALUES (?,?)", userId, feedback.Feedback)

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		response.Data = feedback
		SendResponse(w, response.Status, response)
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		SendResponse(w, response.Status, response)
	}

}

func GetFeedback(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var response model.FeedbacksResponse
	var feedback model.Feedback
	var feedbacks []model.Feedback
	_, userId, _, _ := validateTokenFromCookies(r)

	err := r.ParseForm()
	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		SendResponse(w, response.Status, response)
		return
	}
	// convertedString := strconv.Itoa(userId)

	// query := "SELECT * FROM  Feedbacks  WHERE User_id  = " + convertedString
	rows, err := db.Query("SELECT * FROM  Feedbacks  WHERE User_id  = ? ", userId)

	if err != nil {
		response.Status = 400
		response.Message = "Error Query"
		SendResponse(w, response.Status, response)
		return
	}

	for rows.Next() {
		if err := rows.Scan(&feedback.ID, &feedback.UserId, &feedback.Feedback, &feedback.Date); err != nil {
		} else {
			feedbacks = append(feedbacks, feedback)
		}
	}

	if len(feedbacks) != 0 {
		response.Status = 200
		response.Message = "Succes Get Data"
		response.Data = feedbacks
		SendResponse(w, response.Status, response)
	} else if response.Message == "" {
		response.Status = 400
		response.Message = "Data Not Found"
		SendResponse(w, response.Status, response)
	}

}

func GetAllFeedbacks(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	var response model.FeedbacksResponse
	defer db.Close()

	query := "SELECT * FROM feedbacks"
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

	var feedback model.Feedback
	var feedbacks []model.Feedback

	for rows.Next() {
		if err := rows.Scan(&feedback.ID, &feedback.UserId, &feedback.Feedback, &feedback.Date); err != nil {
			log.Println(err.Error())
		} else {
			feedbacks = append(feedbacks, feedback)
		}
	}

	if len(feedbacks) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = feedbacks
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}

	SendResponse(w, response.Status, response)
	return
}
