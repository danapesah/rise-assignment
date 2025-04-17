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
	router.PUT("/contacts", editContact)
	router.POST("/contacts", createContact)

	router.Run(":8080")
}

type Contact struct {
	ID          int    `json:"id" binding:"required" bson:"_id"`
	FirstName   string `json:"first_name" binding:"required" bson:"first_name"`
	LastName    string `json:"last_name" binding:"required" bson:"last_name"`
	PhoneNumber string `json:"phone_number" binding:"required" bson:"phone_number"`
	Address     string `json:"address" binding:"required" bson:"address"`
}

func getContacts(c *gin.Context) {
	database := getDatabase()
	defer database.disconnect()

	pageStr := c.Query("page")
	if pageStr == "" {
		c.JSON(400, gin.H{"error": "Please provide page number (0+)"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 0 {
		c.JSON(400, gin.H{"error": "Invalid Page Number"})
		return
	}

	contacts := database.loadContactsByPagination("contacts", page)

	c.JSON(200, gin.H{"contacts": contacts})
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

	var contact Contact
	result := database.loadByID(id, "contacts", contact)

	c.JSON(200, gin.H{"contacts": result})
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

	database.delete(id, "contacts")
}

func editContact(c *gin.Context) {
	database := getDatabase()
	defer database.disconnect()

	var contact Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.replace(contact.ID, contact, "contacts")
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
