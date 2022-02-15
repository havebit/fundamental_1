package todo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID    int
	Title string `json:"text" binding:"required"`
	Done  bool
}

var index int = 0
var database = map[int]Todo{}

func NewTaskTodoHandler(c *gin.Context) {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		return
	}

	index++
	todo.ID = index
	database[index] = todo

	fmt.Println(database)
	c.Status(http.StatusCreated)
}
