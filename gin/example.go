package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// go run example.go
// curl http://localhost:3000/ping
// curl http://localhost:3000/user/John
// curl -d "message=help" -X POST http://localhost:3000/form
// curl http://localhost:3000/someXML
// curl http://localhost:3000/someYAML
// curl http://localhost:3000/resolv
// curl -v http://localhost:3000/cookie
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
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	r.StaticFile("/resolv", "/etc/resolv.conf")
	r.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
