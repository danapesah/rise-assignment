package main

import (
	"github.com/gin-gonic/gin"
	"riseAssignment/api"
)

func main() {
	router := gin.Default()

	// Serve static files under /static/
	router.Static("/static", "./static")

	// Serve index.html on /
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	router.GET("/contacts", api.GetContacts)
	router.DELETE("/contacts/:id", api.DeleteContact)
	router.PUT("/contacts", api.EditContact)
	router.POST("/contacts", api.CreateContact)

	router.Run(":8080")
}
