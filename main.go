package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"riseAssignment/api"
)

func main() {
	router := gin.Default()

	router.GET("/contacts", api.GetContacts)
	router.GET("/contacts/:id", api.GetContact)
	router.DELETE("/contacts/:id", api.DeleteContact)
	router.PUT("/contacts", api.EditContact)
	router.POST("/contacts", api.CreateContact)

	router.Run(":8080")
}
