package main

import (
	"github.com/gin-gonic/gin"
	"riseAssignment/api"
)

func main() {
	router := gin.Default()
	api.RegisterMetrics()

	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	router.GET("/contacts", api.GetContacts)
	router.DELETE("/contacts/:id", api.DeleteContact)
	router.PUT("/contacts", api.EditContact)
	router.POST("/contacts", api.CreateContact)

	router.GET("/metrics", api.PrometheusHandler())

	router.Run(":8080")
}
