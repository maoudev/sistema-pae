package student

import (
	"errors"
	"fmt"
	"myproject/internal/pkg/domain"
	"myproject/internal/pkg/ports"
	"time"

	"github.com/xuri/excelize/v2"
)

var (
	ErrStudentNotFound error = errors.New("error: usuario inexistente")
)

type studentService struct {
	repo ports.StudentRepository
}

func NewService(repo ports.StudentRepository) *studentService {
	return &studentService{
		repo: repo,
	}
}

func (u *studentService) CreateStudent(student *domain.Student) error {
	if err := u.repo.CreateStudent(student); err != nil {
		return err
	}

	return nil
}

func (u *studentService) GetStudentByRutWithDash(rut string) (*domain.Student, error) {
	// todo

	student, err := u.repo.GetStudentByRutWithDash(rut)
	if err != nil {
		return nil, err
	}

	if student.Id == 0 || student == nil {
		return nil, ErrStudentNotFound
	}

	return student, nil
}

func (u *studentService) GetStudentByRut(rut string) (*domain.Student, error) {
	// todo

	student, err := u.repo.GetStudentByRut(rut)
	if err != nil {
		return nil, err
	}

	if student.Id == 0 || student == nil {
		return nil, ErrStudentNotFound
	}

	return student, nil
}

func (u *studentService) GetStudentsDocument() (*excelize.File, error) {
	file := excelize.NewFile()

	fechaActual := time.Now().Format("2006_01_02")
	sheetName := fmt.Sprintf("%s_PAE", fechaActual)
	index, err := file.NewSheet(sheetName)
	if err != nil {
		return nil, err
	}

	students, err := u.repo.GetStudentsWhoWentToLunch()
	if err != nil {
		return nil, err
	}

	style, err := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Color: "000000",
		},
		Fill: excelize.Fill{
			Type:  "color",
			Color: []string{"DDDDDD"},
		},
	})
	if err != nil {
		return nil, err
	}

	if err := file.SetCellStyle(sheetName, "A1", "F1", style); err != nil {
		return nil, err
	}

	file.SetColWidth(sheetName, "A", "A", 20)
	file.SetColWidth(sheetName, "B", "B", 40)
	file.SetColWidth(sheetName, "C", "C", 10)
	file.SetColWidth(sheetName, "D", "D", 10)
	file.SetColWidth(sheetName, "E", "E", 20)
	file.SetColWidth(sheetName, "F", "F", 10)

	file.SetCellValue(sheetName, "A1", "Rut")
	file.SetCellValue(sheetName, "B1", "Nombre Completo")
	file.SetCellValue(sheetName, "C1", "Nivel")
	file.SetCellValue(sheetName, "D1", "Letra")
	file.SetCellValue(sheetName, "E1", "Fecha")
	file.SetCellValue(sheetName, "F1", "Hora")

	file.SetActiveSheet(index)

	for i, student := range students {
		row := i + 2 // Empezar en la fila 3 (para dejar espacio para los encabezados)
		file.SetCellValue(sheetName, fmt.Sprintf("A%d", row), student.Rut)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), student.NombreCompleto)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", row), student.Nivel)
		file.SetCellValue(sheetName, fmt.Sprintf("D%d", row), student.Letra)
		file.SetCellValue(sheetName, fmt.Sprintf("E%d", row), student.Fecha)
		file.SetCellValue(sheetName, fmt.Sprintf("F%d", row), student.Hora)
	}

	return file, nil

}

func (u *studentService) InsertStudentWithoutBenefit() error {
	return u.repo.InsertStudentWithoutBenefit()
}
