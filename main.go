package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/contacts", getContacts)
	router.GET("/contacts/:id", getContact)
	router.DELETE("/contacts/:id", deleteContact)
	router.PUT("/contacts/:id", editContact)
	router.POST("/contacts", createContact)

	router.Run("localhost:8080")
}

type Contact struct {
	ID          int    `json:"id" bson:"_id"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Address     string `json:"address" bson:"address"`
}

func getContacts(c *gin.Context) {
	database := getDatabase()
	defer database.disconnect()

	database.load(0, "contacts")
}

func getContact(c *gin.Context) {
	database := getDatabase()
	defer database.disconnect()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	database.load(id, "contacts")
}

func deleteContact(c *gin.Context) {
	database := getDatabase()
	defer database.disconnect()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	doc := Contact{ID: id}
	database.delete(doc, "contacts")
}

func editContact(c *gin.Context) {
	database := getDatabase()
	defer database.disconnect()

	var contact Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.replace(contact, "contacts")
}

func createContact(c *gin.Context) {
	database := getDatabase()
	defer database.disconnect()

	var contact Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.save(contact, "contacts")
}
