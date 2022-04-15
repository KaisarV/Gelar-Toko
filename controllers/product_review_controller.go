package controllers

import (
	"log"
	"net/http"
	"strconv"

	config "GelarToko/config"
	model "GelarToko/models"

	"github.com/gorilla/mux"
)

func GetAllMyProductReviews(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	var response model.ProductReviewsResponse
	defer db.Close()

	_, userId, _, _ := validateTokenFromCookies(r)

	rows, err := db.Query("SELECT * FROM product_reviews WHERE User_Id = ?", userId)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		SendResponse(w, response.Status, response)
		return
	}

	var review model.ProductReview
	var reviews []model.ProductReview

	for rows.Next() {
		if err := rows.Scan(&review.ID, &review.UserId, &review.ProductId, &review.Review, &review.Rating, &review.Date); err != nil {
			log.Println(err.Error())
		} else {
			reviews = append(reviews, review)
		}
	}

	if len(reviews) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = reviews
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	SendResponse(w, response.Status, response)
}

func DeleteMyProductReview(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var response model.ErrorResponse
	_, userId, _, _ := validateTokenFromCookies(r)
	vars := mux.Vars(r)
	productId := vars["productid"]

	query, errQuery := db.Exec(`DELETE FROM product_reviews WHERE User_Id = ? AND Product_Id = ?;`, userId, productId)
	RowsAffected, _ := query.RowsAffected()

	if RowsAffected == 0 {
		response.Status = 400
		response.Message = "store not found"
		SendResponse(w, response.Status, response)
		return
	}

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success Delete Data"
	} else {
		response.Status = 400
		response.Message = "Error Delete Data"
		log.Println(errQuery.Error())
	}

	SendResponse(w, response.Status, response)
}

// // InsertMyProductReview godoc
// // @Summary delete prodduct review.
// // @Description leave a review on the product that has been purchased.
// // @Tags Reviews
// // @Produce json
// // @Param productid path string true "productid"
// // @Param Body body ProductReviewInput true "review's data"
// // @Success 200 {object} model.ProductReviewResponse
// // @Router /review/{productid} [POST]
func InsertMyProductReview(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()

	var review model.ProductReview
	var response model.ProductReviewResponse
	_, userId, _, _ := validateTokenFromCookies(r)
	vars := mux.Vars(r)
	productId := vars["productid"]

	err := r.ParseForm()
	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		return
	}

	// var input ProductReviewInput
	// if c.Request.Header.Get("Content-Type") == "application/json" {
	// 	if err := c.ShouldBindJSON(&input); err != nil {
	// 		response.Status = 400
	// 		response.Message = err.Error()
	// 		c.Header("Content-Type", "application/json")
	// 		c.JSON(http.StatusOK, response)
	// 		return
	// 	}
	// } else {
	review.Review = r.Form.Get("review")
	review.Rating, _ = strconv.Atoi(r.Form.Get("rating"))
	// }

	rows, err := db.Query("SELECT * FROM transactions WHERE User_Id = ? AND Product_Id = ?", userId, productId)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		SendResponse(w, response.Status, response)
		return
	}

	i := 0
	for rows.Next() {
		i++
	}

	if i == 0 {
		response.Status = 400
		response.Message = "you haven't bought this product"
		SendResponse(w, response.Status, response)
		return
	}

	rows, err = db.Query("SELECT * FROM product_reviews WHERE User_Id = ? AND Product_Id = ?", userId, productId)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		SendResponse(w, response.Status, response)
		return
	}

	i = 0
	for rows.Next() {
		i++
	}

	if i != 0 {
		response.Status = 400
		response.Message = "you already review this product"
		SendResponse(w, response.Status, response)
		return
	}

	// review.Review = input.Review
	// review.Rating = input.Rating

	if review.Rating > 5 {
		response.Status = 400
		response.Message = "rating can't be more than 5"
		SendResponse(w, response.Status, response)
		return
	}

	if review.Review == "" {
		response.Status = 400
		response.Message = "Please Insert your review"
		SendResponse(w, response.Status, response)
		return
	}

	if review.Rating == 0 {
		response.Status = 400
		response.Message = "Please insert product's rating"
		SendResponse(w, response.Status, response)
		return
	}

	res, errQuery := db.Exec("INSERT INTO product_reviews(User_Id, Product_Id,  Review, Rating) VALUES(?, ?, ?, ?)", userId, productId, review.Review, review.Rating)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		review.ID = int(id)
		response.Data = review
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}
	SendResponse(w, response.Status, response)

}

func UpdateMyProductReview(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()

	var review model.ProductReview
	var response model.ProductReviewResponse
	_, userId, _, _ := validateTokenFromCookies(r)
	vars := mux.Vars(r)
	productId := vars["productid"]

	// var input ProductReviewInput
	err := r.ParseForm()
	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		return
	}

	// if c.Request.Header.Get("Content-Type") == "application/json" {
	// 	if err := c.ShouldBindJSON(&input); err != nil {
	// 		response.Status = 400
	// 		response.Message = err.Error()
	// 		c.Header("Content-Type", "application/json")
	// 		c.JSON(http.StatusOK, response)
	// 		return
	// 	}
	// } else {
	// input.Review = c.PostForm("review")
	// input.Rating, _ = strconv.Atoi(c.PostForm("rating"))
	// }

	review.Review = r.Form.Get("review")
	review.Rating, _ = strconv.Atoi(r.Form.Get("rating"))

	rows, _ := db.Query("SELECT Review, Rating FROM product_reviews WHERE User_Id = ? AND Product_Id = ?", userId, productId)
	var prevDatas []model.ProductReview
	var prevData model.ProductReview

	for rows.Next() {
		if err := rows.Scan(&prevData.Review, &prevData.Rating); err != nil {
			log.Println(err.Error())
		} else {
			prevDatas = append(prevDatas, prevData)
		}
	}

	if len(prevDatas) > 0 {
		if review.Review == "" {
			review.Review = prevDatas[0].Review
		}
		if review.Rating == 0 {
			review.Rating = prevDatas[0].Rating
		}

		_, errQuery := db.Exec(`UPDATE product_reviews SET Review = ?, Rating = ? WHERE User_Id = ? AND Product_Id = ?`, review.Review, review.Rating, userId, productId)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Success Update Data"
			review.ID = prevData.ID
			response.Data = review
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
