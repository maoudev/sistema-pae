package ports

import (
	"myproject/internal/pkg/domain"

	"github.com/xuri/excelize/v2"
)

type StudentRepository interface {
	CreateStudent(student *domain.Student) error
	GetStudentByRut(rut string) (*domain.Student, error)
	GetStudentByRutWithDash(rut string) (*domain.Student, error)
	GetStudentsWhoWentToLunch() ([]*domain.StudentsRequest, error)
	InsertStudentWithoutBenefit() error
}

type StudentService interface {
	CreateStudent(student *domain.Student) error
	GetStudentByRut(rut string) (*domain.Student, error)
	GetStudentByRutWithDash(rut string) (*domain.Student, error)
	GetStudentsDocument() (*excelize.File, error)
	InsertStudentWithoutBenefit() error
}
