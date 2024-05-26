package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Todo struct {
    ID          uint   json:`"id"`
    Title       string json:"title"
    Description string json:"description"
    Completed   bool   json:"completed"
}

var todos []Todo
var nextID uint = 1

func main() {
    router := gin.Default()

    router.Static("/static", "./static")
    router.GET("/", func(c *gin.Context) {
        c.File("./static/index.html")
    })

    router.GET("/todos", getTodos)
    router.POST("/todos", createTodo)
    router.GET("/todos/:id", getTodo)
    router.PUT("/todos/:id", updateTodo)
    router.DELETE("/todos/:id", deleteTodo)

    router.Run(":8080")
}

func getTodos(c *gin.Context) {
    c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
    var todo Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    todo.ID = nextID
    nextID++
    todos = append(todos, todo)
    c.Status(http.StatusCreated)
}

func getTodo(c *gin.Context) {
    id := c.Param("id")
    for _, todo := range todos {
        if c.Param("id") == string(rune(todo.ID)) {
            c.JSON(http.StatusOK, todo)
            return
        }
    }
    c.Status(http.StatusNotFound)
}

func updateTodo(c *gin.Context) {
    id := c.Param("id")
    var updatedTodo Todo
    if err := c.ShouldBindJSON(&updatedTodo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    for i, todo := range todos {
        if c.Param("id") == string(rune(todo.ID)) {
            todos[i] = updatedTodo
            todos[i].ID = todo.ID
            c.Status(http.StatusNoContent)
            return
        }
    }
    c.Status(http.StatusNotFound)
}

func deleteTodo(c *gin.Context) {
    id := c.Param("id")
    for i, todo := range todos {
        if c.Param("id") == string(rune(todo.ID)) {
            todos = append(todos[:i], todos[i+1:]...)
            c.Status(http.StatusNoContent)
            return
        }
    }
    c.Status(http.StatusNotFound)
}