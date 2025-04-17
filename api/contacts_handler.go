package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"riseAssignment/db"
	"strconv"
	"time"
)

type Contact struct {
	ID          int    `json:"id" binding:"required" bson:"_id"`
	FirstName   string `json:"first_name" binding:"required" bson:"first_name"`
	LastName    string `json:"last_name" binding:"required" bson:"last_name"`
	PhoneNumber string `json:"phone_number" binding:"required" bson:"phone_number"`
	Address     string `json:"address" binding:"required" bson:"address"`
}

func GetContacts(c *gin.Context) {
	database := db.GetDatabase()
	defer database.Disconnect()

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

	cursor, err := database.LoadByPagination("contacts", page)

	var contacts []Contact
	if err = cursor.All(context.TODO(), &contacts); err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{"contacts": contacts})
}

func GetContact(c *gin.Context) {
	database := db.GetDatabase()
	defer database.Disconnect()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var contact Contact
	result := database.LoadByID(id, "contacts", contact)

	c.JSON(200, gin.H{"contacts": result})
}

func DeleteContact(c *gin.Context) {
	database := db.GetDatabase()

	defer database.Disconnect()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	database.Delete(id, "contacts")
}

func EditContact(c *gin.Context) {
	database := db.GetDatabase()

	defer database.Disconnect()

	var contact Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.Replace(contact.ID, contact, "contacts")
}

func CreateContact(c *gin.Context) {
	database := db.GetDatabase()

	defer database.Disconnect()

	var contact Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.Save(contact, "contacts")
}
