package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllTodos godoc
// @Summary get all items in the todo list
// @ID get-all-todos
// @Produce json
// @Success 200 {object} todo
// @Router /todo [get]
func GetAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todoList)
}

// GetTodoByID godoc
// @Summary get a todo item by ID
// @ID get-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} todo
// @Failure 404 {object} message
// @Router /todo/{id} [get]
func GetTodoByID(c *gin.Context) {
	ID := c.Param("id")

	// loop through todoList and return item with matching ID
	for _, todo := range todoList {
		if todo.ID == ID {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	// return error message if todo is not found
	r := message{"todo not found"}
	c.JSON(http.StatusNotFound, r)
}

// CreateTodo godoc
// @Summary add a new item to the todo list
// @ID create-todo
// @Produce json
// @Param data body todo true "todo data"
// @Success 200 {object} todo
// @Failure 400 {object} message
// @Router /todo [post]
func CreateTodo(c *gin.Context) {
	var newTodo todo

	// bind the received JSON data to newTodo
	if err := c.BindJSON(&newTodo); err != nil {
		r := message{"an error occurred while creating todo"}
		c.JSON(http.StatusBadRequest, r)
		return
	}

	// add the new todo item to todoList
	todoList = append(todoList, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

// DeleteTodo godoc
// @Summary delete a todo item by ID
// @ID delete-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} todo
// @Failure 404 {object} message
// @Router /todo/{id} [delete]
func DeleteTodo(c *gin.Context) {
	ID := c.Param("id")

	// loop through todoList and delete item with matching ID
	for index, todo := range todoList {
		if todo.ID == ID {
			todoList = append(todoList[:index], todoList[index+1:]...)
			r := message{"successfully deleted todo"}
			c.JSON(http.StatusOK, r)
			return
		}
	}

	// return error message if todo is not found
	r := message{"todo not found"}
	c.JSON(http.StatusNotFound, r)
}
