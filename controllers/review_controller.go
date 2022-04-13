package controllers

import (
	config "GelarToko/config"
	model "GelarToko/models"
	"log"
	"net/http"
	"strconv"
)

func InsertNewReviews(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var response model.ProductReviewResponse
	var review model.ProductReview
	_, userId, _, _ := validateTokenFromCookies(r)

	err := r.ParseForm()
	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		SendResponse(w, response.Status, response)
		return
	}

	review.UserId, _ = strconv.Atoi(r.Form.Get("userID"))
	review.Review = r.Form.Get("Review")

	_, errQuery := db.Exec("SELECT FROM transactions WHERE ID = ?'", review.UserId)
	if errQuery != nil {
		db.Exec("INSERT INTO Reviews (userid, review, date) VALUES (?,?,SYSDATE)",
			userId,
			review.Review,
		)
	}

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		response.Data = review
		SendResponse(w, response.Status, response)
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		SendResponse(w, response.Status, response)
	}

}

func GetReviews(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var response model.ProductReviewResponse
	var review model.ProductReview
	var reviews []model.ProductReview
	_, userId, _, _ := validateTokenFromCookies(r)

	err := r.ParseForm()
	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		SendResponse(w, response.Status, response)
		return
	}
	convertedString := strconv.Itoa(userId)
	query := ""

	_, errQuery := db.Exec("SELECT FROM transactions WHERE ID = ?'", review.UserId)
	if errQuery != nil {
		query = "SELECT * FROM  Reviews  WHERE User_id  = " + convertedString
	}
	rows, err := db.Query(query)

	if err != nil {
		response.Status = 400
		response.Message = "Error Query"
		SendResponse(w, response.Status, response)
		return
	}

	for rows.Next() {
		if err := rows.Scan(&review.ID, &review.UserId, &review.Review, &review.Date); err != nil {
		} else {
			reviews = append(reviews, review)
		}
	}

	if len(reviews) != 0 {
		response.Status = 200
		response.Message = "Succes Get Data"
		response.Data = review
		SendResponse(w, response.Status, response)
	} else if response.Message == "" {
		response.Status = 400
		response.Message = "Data Not Found"
		SendResponse(w, response.Status, response)
	}

}

func GetAllReviews(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	var review model.ProductReview
	var reviews []model.ProductReview
	var response model.ProductReviewResponse
	defer db.Close()

	query := ""
	_, errQuery := db.Exec("SELECT FROM transactions WHERE ID = ?'", review.UserId)
	if errQuery != nil {
		query = "SELECT * FROM Reviews"
		id := r.URL.Query()["id"]
		if id != nil {
			query += " WHERE id = " + id[0]
		}
	}

	rows, err := db.Query(query)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		SendResponse(w, response.Status, response)
		return
	}

	for rows.Next() {
		if err := rows.Scan(&review.ID, &review.UserId, &review.Review, &review.Date); err != nil {
			log.Println(err.Error())
		} else {
			reviews = append(reviews, review)
		}
	}

	if len(reviews) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = review
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}

	SendResponse(w, response.Status, response)
	return
}
