package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Author  string `json:"Author"`
	Ordered bool   `json:"ordered"`
}

var books = []Book{
	{ID: 1, Name: "Sergeant Keroro", Author: "Keroro Gunsō", Ordered: false},
	{ID: 2, Name: "Naruto", Author: "Kishimoto Masashi", Ordered: false},
	{ID: 3, Name: "Dragon Ball", Author: "Akira Toriyama", Ordered: false},
	{ID: 4, Name: "One Piece", Author: "Oda Eiichirō", Ordered: true},
}

func getBooks(c *gin.Context) {
	nameQuery := c.Query("name")

	filter := []Book{}
	for _, d := range books {
		if nameQuery == "" || strings.Contains(strings.ToLower(d.Name), strings.ToLower(nameQuery)) {
			filter = append(filter, d)
		}
	}

	c.JSON(http.StatusOK, filter)
}

func orderBook(c *gin.Context) {
	var req struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, j := range books {
		if j.ID == req.ID {
			if j.Ordered {
				c.JSON(http.StatusConflict, gin.H{"message": "TEST"})
				return
			}
			books[i].Ordered = true
			c.JSON(http.StatusOK, books[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
}

func main() {
	r := gin.Default()

	// health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "healthy"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/books", getBooks)
		api.POST("/order", orderBook)
	}

	r.Run(":8080")
}
