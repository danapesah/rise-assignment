package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"riseAssignment/db"
	"strconv"
)

type Contact struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName   string             `json:"first_name" binding:"required" bson:"first_name,omitempty"`
	LastName    string             `json:"last_name" binding:"required" bson:"last_name,omitempty"`
	PhoneNumber string             `json:"phone_number" binding:"required" bson:"phone_number,omitempty"`
	Address     string             `json:"address" binding:"required" bson:"address,omitempty"`
}

func GetContacts(c *gin.Context) {
	database := db.GetDatabase()
	defer database.Disconnect()

	filterContact := parseFilerParams(c)
	var contacts []Contact

	if filterContact == (Contact{}) {
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

		cursor, _ := database.LoadByPagination("contacts", page)
		if err := cursor.All(context.TODO(), &contacts); err != nil {
			panic(err)
		}
	} else {
		cursor, _ := database.Load(filterContact, "contacts")
		if err := cursor.All(context.TODO(), &contacts); err != nil {
			panic(err)
		}
	}

	c.JSON(200, gin.H{"contacts": contacts})
}

func parseFilerParams(c *gin.Context) (filterContact Contact) {
	filterParam := c.Query("first_name")
	if filterParam != "" {
		filterContact.FirstName = filterParam
	}
	filterParam = c.Query("last_name")
	if filterParam != "" {
		filterContact.LastName = filterParam
	}
	filterParam = c.Query("phone_number")
	if filterParam != "" {
		filterContact.PhoneNumber = filterParam
	}
	filterParam = c.Query("address")
	if filterParam != "" {
		filterContact.Address = filterParam
	}

	return filterContact
}

func DeleteContact(c *gin.Context) {
	database := db.GetDatabase()

	defer database.Disconnect()

	idStr := c.Param("id")
	var contact Contact
	contact.ID, _ = primitive.ObjectIDFromHex(idStr)

	database.Delete(contact, "contacts")
}

func EditContact(c *gin.Context) {
	database := db.GetDatabase()
	defer database.Disconnect()

	var contact Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	filterContact := Contact{ID: contact.ID}

	database.Replace(filterContact, contact, "contacts")
}

func CreateContact(c *gin.Context) {
	database := db.GetDatabase()
	defer database.Disconnect()

	var contact Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	contact.ID = primitive.NewObjectID()

	database.Save(contact, "contacts")
}
