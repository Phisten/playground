package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {

	// gin
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to my website!")
	})

	r.GET("/books/:title/page/:page", func(c *gin.Context) {
		title := c.Param("title")
		page := c.Param("page")
		c.String(200, "You've requested the book: %s on page %s\n", title, page)
	})
	go func() {
		r.Run(":3500")
	}()

	// 第一個伺服器，監聽在 8080 port
	router1 := gin.Default()
	router1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Server 1!",
		})
	})
	go func() {
		router1.Run(":3510")
	}()

	select {}
}
