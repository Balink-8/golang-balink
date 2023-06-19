package main

import (
    "golang-img-api/controllers" //add this
    "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Golang Img RESTFull API"})
	})
	// Upload Image
    router.POST("/upload", controllers.FileUpload())

	router.Run("localhost:5000")
}