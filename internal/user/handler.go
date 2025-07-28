package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	u := r.Group("/users")
	{
		u.GET("/", GetAll)
		u.POST("/", CreateHandler)
		u.PUT("/:id/finish", Finish)
	}
}

func GetAll(c *gin.Context) {
	users, err := FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar usuários"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreateHandler(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := CreateUser(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}
	c.JSON(http.StatusCreated, u)
}

func Finish(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := FinishUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao finalizar usuário"})
		return
	}
	c.Status(http.StatusNoContent)
}
