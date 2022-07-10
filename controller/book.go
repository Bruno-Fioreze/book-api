package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/database"
	"github.com/hyperyuri/webapi-with-go/models"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer.",
		})
		return
	}

	db := database.GetDatabase()
	var book models.Book
	err = db.First(&book, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book" + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()
	var book models.Book
	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book" + err.Error(),
		})
		return
	}

}
