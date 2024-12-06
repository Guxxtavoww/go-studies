package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{Id: "1", Title: "Book 1", Author: "Gugu", Quantity: 14},
	{Id: "2", Title: "Book 2", Author: "Vaga", Quantity: 15},
	{Id: "3", Title: "Book 3", Author: "Tee", Quantity: 12},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func addBook(context *gin.Context) {
	var newBook Book

	if err := context.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	context.IndentedJSON(http.StatusCreated, books)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/books", addBook)

	router.Run("localhost:3333")
}
