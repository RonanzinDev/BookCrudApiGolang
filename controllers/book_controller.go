package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ronanzindev/book-api-golang/database"
	"github.com/ronanzindev/book-api-golang/models"
)

// Pegando um book pelo id dele
func ShowBook(c *gin.Context) {
	id := c.Param("id") // pegando id dos paremetros
	// convertend o id para int e veficando se tem algum error
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id have to be integer",
		})
		return
	}
	// chamdando o banco de dados
	db := database.GetDataBase()
	// instaciando o book
	var book models.Book
	// verificando se o book existe, caso não exista retorna um error
	err = db.First(&book, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDataBase()

	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind Json " + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book" + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDataBase()
	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "cannot list books " + err.Error()})
	}
	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDataBase()

	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind Json " + err.Error(),
		})
		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create update" + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id") // pegando id dos paremetros
	// convertend o id para int e veficando se tem algum error
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id have to be integer",
		})
		return
	}

	db := database.GetDataBase()
	err = db.Delete(&models.Book{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete de book" + err.Error(),
		})
	}

	c.JSON(204, gin.H{
		"deleteStatus": "delete sucess",
	})

}
