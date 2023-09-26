package main

import (
	"log"
	. "github.com/farimarwat/go-books/database"
	. "github.com/farimarwat/go-books/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {	
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	
	defer Disconnect()
	router := gin.Default()
	router.POST("/books",CreateBook)
	router.GET("/books",ListBooks)
	router.GET("/books/:name",FindBook)
	router.Run("localhost:8080")
}

