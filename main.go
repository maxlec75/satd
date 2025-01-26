package main

import (
	"SATD/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	controllers.InitDatabase()
	r := gin.Default()
	r.GET("/books", controllers.FindBooks)
	r.Run()
}
