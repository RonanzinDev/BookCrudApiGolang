package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ronanzindev/book-api-golang/database"
	"github.com/ronanzindev/book-api-golang/models"
	"github.com/ronanzindev/book-api-golang/services"
)

func CreateUser(c *gin.Context) {
	db := database.GetDataBase()

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind Json " + err.Error(),
		})
		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user" + err.Error(),
		})
		return
	}

	c.Status(2045)
}
