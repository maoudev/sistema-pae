package domain

type Student struct {
	Id             int    `json:"id" db:"id"`
	Rut            string `json:"rut" db:"rut" binding:"required"`
	Nombre         string `json:"nombre" db:"nombre" binding:"required"`
	Nivel          int    `json:"nivel" db:"nivel" binding:"required"`
	Letra          string `json:"letra" db:"letra" binding:"required"`
	Almorzo        bool   `json:"almorzo"`
}
