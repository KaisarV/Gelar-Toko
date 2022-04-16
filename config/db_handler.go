package controllers

import (
	"GelarToko/utils"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db_name := utils.Getenv("DATABASE_NAME", "1")
	db_connection := utils.Getenv("DATABASE_CONNECTION", "1")
	db_username := utils.Getenv("DATABASE_USERNAME", "1")
	db_password := utils.Getenv("DATABASE_PASSWORD", "1")
	db_port := utils.Getenv("DATABASE_PORT", "1")
	db_host := utils.Getenv("DATABASE_HOST", "1")

	db, err := sql.Open(db_connection, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_username, db_password, db_host, db_port, db_name))

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Success Connect DB")
	}

	return db

}
