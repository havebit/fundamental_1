package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title string `json:"text" binding:"required"`
	Done  bool
}

func (Todo) TableName() string {
	return "todos"
}

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) NewTaskTodoHandler(c *gin.Context) {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		return
	}

	// h.db.Find()
	if err := h.db.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) ListTaskTodoHandler(c *gin.Context) {
}
