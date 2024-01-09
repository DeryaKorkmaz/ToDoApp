package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var todos []string

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Todos": todos,
		})
	})

	api := router.Group("/api")
	{
		api.GET("/todos", getTodos)
		api.POST("/todos", addTodo)
		api.PUT("/todos/:id", updateTodo)
		api.DELETE("/todos/:id", deleteTodo)
	}

	router.Run(":8080")
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func addTodo(c *gin.Context) {
	var todo struct {
		Text string `json:"text"`
	}
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todos = append(todos, todo.Text)
	c.JSON(http.StatusOK, gin.H{"message": "Todo added successfully"})
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo struct {
		Text string `json:"text"`
	}
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	index := getIndexByID(id)
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	todos[index] = todo.Text
	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	index := getIndexByID(id)
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	todos = append(todos[:index], todos[index+1:]...)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func getIndexByID(id string) int {
	for i, todo := range todos {
		if todo == id {
			return i
		}
	}
	return -1
}
