package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json: "id"`
	Title    string `json: "title"`
	Author   string `json: "author"`
	Quantity int    `json: "quantity"`
}

var library = []book{
	{ID: "1", Title: "Refactoring", Author: "Martin Fowler", Quantity: 2},
	{ID: "2", Title: "Structure and Interpretation of Computer Programs", Author: "Harold Abelson", Quantity: 5},
	{ID: "3", Title: "Extreme Programming Explained", Author: "Kent Beck", Quantity: 8},
	{ID: "4", Title: "System Design Interview", Author: "Alex Xu", Quantity: 1},
	{ID: "5", Title: "Microservices Patterns", Author: "Chris Richardson", Quantity: 3},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, library)
}

func createBook(context *gin.Context) {
	var newBook book

	if err := context.BindJSON(&newBook); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	library = append(library, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/books", createBook)

	router.Run("localhost:8080")
}
