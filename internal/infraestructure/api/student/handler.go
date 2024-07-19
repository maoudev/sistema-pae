package student

import (
	"errors"
	"fmt"
	"myproject/internal/infraestructure/repositories/mysql"
	"myproject/internal/pkg/domain"
	"myproject/internal/pkg/ports"
	stdnt "myproject/internal/pkg/services/student"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type studentHandler struct {
	studentService ports.StudentService
}

func newHandler(service ports.StudentService) *studentHandler {
	return &studentHandler{
		studentService: service,
	}
}

func (u *studentHandler) CreateStudent(c *gin.Context) {
	student := &domain.Student{}

	if err := c.BindJSON(student); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := u.studentService.CreateStudent(student); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return

	}

	c.JSON(http.StatusOK, nil)
}

func (u *studentHandler) GetStudent(c *gin.Context) {
	rut := c.Param("rut")

	student, err := u.studentService.GetStudentByRut(rut)
	if err != nil {
		if errors.Is(err, mysql.ErrStudentNotFound) || errors.Is(err, stdnt.ErrStudentNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		} else {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	}

	c.JSON(http.StatusOK, student)
}

func (u *studentHandler) GetStudentWithDash(c *gin.Context) {
	rut := c.Param("rut")

	student, err := u.studentService.GetStudentByRutWithDash(rut)
	if err != nil {
		if errors.Is(err, mysql.ErrStudentNotFound) || errors.Is(err, stdnt.ErrStudentNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		} else {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	}

	c.JSON(http.StatusOK, student)
}

func (u *studentHandler) GetStudents(c *gin.Context) {
	file, err := u.studentService.GetStudentsDocument()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fechaActual := time.Now().Format("2006_01_02")
	fileName := fmt.Sprintf("%s_PAE.%s", fechaActual, "xlsx")

	file.SaveAs(fileName)
	defer os.Remove(fileName)

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File(fileName)
}

func (u *studentHandler) InsertStudentWithoutBenefit(c *gin.Context) {
	if err := u.studentService.InsertStudentWithoutBenefit(); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
