package api

import (
	"myproject/internal/infraestructure/api/student"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(e *gin.Engine) {
	student.SetupRoutes(e)
}
