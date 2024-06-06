package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"myproject/internal/pkg/domain"
	"time"
)

var (
	ErrStudentNotFound error = errors.New("alumno no encontrado")
)

type client struct {
	db *sql.DB
}

func NewClient() *client {
	return &client{
		db: connect(),
	}
}

func (c *client) CreateStudent(student *domain.Student) error {
	query := "insert into alumnos(rut, nombre, nivel, letra) values(?, ?, ?, ?)"

	_, err := c.db.Exec(query, student.Rut, student.Nombre, student.Nivel, student.Letra)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) GetStudentByRutWithDash(rut string) (*domain.Student, error) {
	student := &domain.Student{}

	query := "select * from alumnos WHERE rut = ?"

	rows, err := c.db.Query(query, rut)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		var rut, nombre, letra string
		var id, nivel int
		err := rows.Scan(&id, &rut, &nombre, &nivel, &letra)
		if err != nil {
			return nil, err
		}

		student = &domain.Student{Id: id, Rut: rut, Nombre: nombre, Nivel: nivel, Letra: letra, Almorzo: false}

		if student.Id != 0 {
			fechaActual := time.Now().Format("2006-01-02")

			query := fmt.Sprintf("SELECT COUNT(*) FROM ingreso_alumnos WHERE alumno_id = ? AND fecha >= '%s 00:00' AND fecha <= '%s 23:59';", fechaActual, fechaActual)

			var cantidad int
			if err := c.db.QueryRow(query, student.Id).Scan(&cantidad); err != nil {
				fmt.Print(err.Error())

				return nil, err
			}

			if cantidad > 0 {
				student.Almorzo = true
			} else {
				if student.Id != 0 {
					query = "insert into ingreso_alumnos(alumno_id, fecha) values(?, NOW())"

					_, err := c.db.Exec(query, student.Id)
					if err != nil {
						return nil, err
					}
				}
				student.Almorzo = false
			}

		} else {
			return nil, ErrStudentNotFound
		}
	}

	return student, nil
}

func (c *client) GetStudentByRut(rut string) (*domain.Student, error) {
	student := &domain.Student{}

	query := "select * from alumnos WHERE LEFT(rut, 8) = ?"

	rows, err := c.db.Query(query, rut)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		var rut, nombre, letra string
		var id, nivel int
		err := rows.Scan(&id, &rut, &nombre, &nivel, &letra)
		if err != nil {
			return nil, err
		}

		student = &domain.Student{Id: id, Rut: rut, Nombre: nombre, Nivel: nivel, Letra: letra, Almorzo: false}

		if student.Id != 0 {
			fechaActual := time.Now().Format("2006-01-02")

			query := fmt.Sprintf("SELECT COUNT(*) FROM ingreso_alumnos WHERE alumno_id = ? AND fecha >= '%s 00:00' AND fecha <= '%s 23:59';", fechaActual, fechaActual)

			var cantidad int
			if err := c.db.QueryRow(query, student.Id).Scan(&cantidad); err != nil {
				return nil, err
			}

			if cantidad > 0 {
				student.Almorzo = true
			} else {
				if student.Id != 0 {
					query = "insert into ingreso_alumnos(alumno_id, fecha) values(?, NOW())"

					_, err := c.db.Exec(query, student.Id)
					if err != nil {
						return nil, err
					}
				}
				student.Almorzo = false
			}

		} else {
			return nil, ErrStudentNotFound
		}
	}

	return student, nil
}

func (c *client) GetStudentsWhoWentToLunch() ([]*domain.StudentsRequest, error) {
	var students []*domain.StudentsRequest

	fechaActual := time.Now().Format("2006-01-02")
	query := `select rut, nombre, nivel, letra, date_format(fecha, "%d-%m-%Y"), DATE_FORMAT(fecha, "%H:%i") from alumnos left join ingreso_alumnos on alumnos.id = ingreso_alumnos.alumno_id where fecha >= '` + fechaActual + " 00:00' and fecha <= '" + fechaActual + " 23:59' order by nivel, letra;"

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		var rut, nombre, letra, fecha, hora string
		var nivel int

		if err := rows.Scan(&rut, &nombre, &nivel, &letra, &fecha, &hora); err != nil {
			return nil, err
		}

		student := &domain.StudentsRequest{Rut: rut, NombreCompleto: nombre, Nivel: nivel, Letra: letra, Fecha: fecha, Hora: hora}

		students = append(students, student)
	}

	return students, nil
}

func (c *client) InsertStudentWithoutBenefit() error {
	query := "insert into ingreso_alumnos(alumno_id, fecha) values((select id from alumnos where rut = '00000000-0'), NOW())"
	_, err := c.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
