package main

import (
	"fmt"
	"log"
	"os"
	"sample-todo-app/server"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//ここからDB情報（ご自分の環境に合わせて変更してください）
	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	dbx, err := sqlx.Connect("postgres", connInfo)
	//ここまで

	router := gin.Default()
	router.LoadHTMLGlob("assets/views/*.html")
	router.Static("/assets/js", "./assets/js")
	router.Static("/assets/css", "./assets/css")

	router.GET("/", server.ShowTodoList(dbx))

	if err := router.Run(":8080"); err != nil {
		log.Fatal("RouterError:", err)
	}
}
