package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "yossef-test",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		firstName := c.DefaultQuery("firstname", "lalala")
		lastName := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s", name, firstName, lastName)
	})

	// r.POST("/form_post", func(c *gin.Context) {
	// 	message := c.PostForm("message")
	// 	nick := c.DefaultPostForm("nick", "qwe")

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": gin.H{
	// 			"status_code": http.StatusOK,
	// 			"status": "ok"
	// 		},
	// 		"message": message,
	// 		"nick": nick,
	// 	})
	// })

	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}