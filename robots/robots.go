package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	todo struct {
		gorm.Model
		Title     string `json:"title"`
		URL       string `json:"url"`
		Completed int    `json:"completed"`
	}
	// transformedTodo represents a formatted todo
	transformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		URL       string `json:"url"`
		Completed bool   `json:"completed"`
	}
)

var db *gorm.DB

func init() {
	dbPasswd, ok := os.LookupEnv("ROBOTS_PASS")
	if !ok {
		panic("must set ROBOTS_PASS in environment")
	}
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "robotron:"+dbPasswd+"@tcp(localhost:3306)/robots?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}
	//Migrate the schema
	db.SingularTable(true)
	db.AutoMigrate(&todo{})
}

func main() {
	router := setupRouter()
	router.Run(":3000") // listen and serve on 0.0.0.0:3000
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/robots/api/v1.0/recipes")
	{
		// v1.POST("/", createRecipe)
		v1.GET("/", fetchAllRecipe)
		// v1.GET("/:id", fetchSingleRecipe)
		//v1.PUT("/:id", updateRecipe)
		// v1.DELETE("/:id", deleteRecipe)
	}

	return router
}

func fetchAllRecipe(c *gin.Context) {
	var todos []todo
	var _todos []transformedTodo

	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	//transforms the todos for building a good response
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, transformedTodo{ID: item.ID, Title: item.Title, URL: item.URL, Completed: completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}
