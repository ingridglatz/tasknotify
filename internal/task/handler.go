package task

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	t := r.Group("/tasks")
	{
		t.GET("/", GetAll)
		t.POST("/", CreateHandler)
		t.PUT("/:id/finish", Finish)
	}
}

func GetAll(c *gin.Context) {
	tasks, err := FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar tarefas"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func CreateHandler(c *gin.Context) {
	var t Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := CreateTask(&t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar tarefa"})
		return
	}
	c.JSON(http.StatusCreated, t)
}

func Finish(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := FinishTask(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao finalizar tarefa"})
		return
	}
	c.Status(http.StatusNoContent)
}
