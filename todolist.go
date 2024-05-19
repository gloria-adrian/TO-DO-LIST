package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var todos []Todo

func main() {
	//Initialize Gin
	router := gin.Default()

	//Define routes
	router.GET("/todos", getTodos)
	router.POST("/todos", createTodo)
	router.GET("/todos/:id", getTodos)
	router.PUT("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)

	//Run the server
	router.Run(":8080")
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	todos = append(todos, todo)
	c.Status(http.StatusCreated)
}
func updateTodo(c *gin.Context) {
	id := c.Param("id")
	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}
	c.Status(http.StatusNotFound)
}

func updatedTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = updatedTodo
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.Status(http.StatusNotFound)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.Status(http.StatusNotFound)
}
