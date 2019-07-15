package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// go run example.go
// curl http://localhost:3000/ping
// curl http://localhost:3000/user/John
// curl -d "message=help" -X POST http://localhost:3000/form
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s\n", name)
	})
	r.POST("/form", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
