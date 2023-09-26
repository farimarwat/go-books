package routes

import (
	"net/http"
	. "github.com/farimarwat/go-books/database"
	. "github.com/farimarwat/go-books/models"
	"github.com/gin-gonic/gin"

)

func ListBooks(c *gin.Context) {
	listbooks := List_Books()
	if listbooks == nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"List is empty"})
		return
	}
	//Response to api
	c.IndentedJSON(http.StatusOK,gin.H{"message":"Data fetched successfully","data":listbooks})
}

func CreateBook(c *gin.Context){
	var book Book
	err := c.BindJSON(&book)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Bad request"})
		return
	}
	id := Create_Book(book)
	//end mongo queries
	c.IndentedJSON(http.StatusCreated,gin.H{"message":"Book Created Successful","data":id})
}

func FindBook(c *gin.Context){
	name := c.Param("name")
	book := Find_Book(name)
	if book == nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found"})
		return
	}
	//end mongo queries
	c.IndentedJSON(http.StatusOK,gin.H{"data":book})
}


