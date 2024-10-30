package main

import (
	"anammaulana/backend-api/controllers"
	"anammaulana/backend-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//inisialiasai Gin
	router := gin.Default()

	//panggil koneksi database
	models.ConnectDatabase()

	//membuat route dengan method GET
	router.GET("/", func(c *gin.Context) {

		//return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	//membuat route get all posts
	router.GET("/api/posts", controllers.FindPosts)

	//membuat route store post
	router.POST("/api/posts", controllers.StorePost)

	//membuat route detail post
	router.GET("/api/posts/:id", controllers.FindPostById)

	//mulai server dengan port 3000
	router.Run(":3000")
}