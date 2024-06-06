package student

import (
	"myproject/internal/infraestructure/repositories/mysql"
	"myproject/internal/pkg/services/student"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(e *gin.Engine) {
	repo := mysql.NewClient()
	service := student.NewService(repo)
	handler := newHandler(service)

	e.POST("/student", handler.CreateStudent)
	e.GET("/student/scanner/:rut", handler.GetStudent)
	e.GET("/student/search/:rut", handler.GetStudentWithDash)
	e.GET("/students", handler.GetStudents)
	e.POST("/student/non-beneficiary", handler.InsertStudentWithoutBenefit)
}
