package routes

import (
	"tasknotify/internal/task"
	"tasknotify/internal/user"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	task.RegisterRoutes(r)
	user.RegisterRoutes(r)
}
