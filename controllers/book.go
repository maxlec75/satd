package controllers

import (
	"SATD/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Library []models.Book
var Counter int

func InitDatabase() {
	Counter = 1

	book1 := models.Book{
		Id:     1,
		Title:  "Titre",
		Author: "Auteur",
	}
	Library = append(Library, book1)

}

func FindBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": Library})
}
