package main

import (
	"errors"
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

func bookById(context *gin.Context) {
	id := context.Param("id")
	book, err := getBookById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"ERROR": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range library {
		if b.ID == id {
			return &library[i], nil
		}
	}

	return nil, errors.New("The book that you're searching isn't on our library")
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

func checkoutBook(context *gin.Context) {
	id, ok := context.GetQuery("id")

	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": "You're missing something on your query"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"ERROR": err.Error()})
		return
	}

	if book.Quantity <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": "This book isn't available"})
		return
	}

	book.Quantity -= 1
	context.IndentedJSON(http.StatusOK, book)
}

func checkinBook(context *gin.Context) {
	id, ok := context.GetQuery("id")

	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": "You're missing something on your query"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"ERROR": err.Error()})
		return
	}

	book.Quantity += 1
	context.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)
	router.POST("/checkout", checkoutBook)
	router.POST("/checkin", checkinBook)

	router.Run("localhost:8080")
}
