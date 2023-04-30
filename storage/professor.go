package storage

import (
	"database/sql"

	"github.com/jaehong21/ga-be/entity"
)

func CreateProfessor(tx *sql.Tx, id int, password string, name string, major string, phone string, email string) (entity.Professor, error) {
	var professor entity.Professor

	row := tx.QueryRow("INSERT INTO \"professor\" (id, password, name, major, phone, email) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *", id, password, name, major, phone, email)
	err := row.Scan(&professor.ID, &professor.Password, &professor.Name, &professor.Major, &professor.Phone, &professor.Email)
	if err != nil {
		return professor, err
	}

	return professor, nil
}

func FindOneProfessor(db *sql.DB, id int) (entity.Professor, error) {
	var professor entity.Professor

	row := db.QueryRow("SELECT id, password, name, major, phone, email FROM \"professor\" WHERE id=$1", id)
	err := row.Scan(&professor.ID, &professor.Password, &professor.Name, &professor.Major, &professor.Phone, &professor.Email)
	if err != nil {
		return professor, err
	}

	return professor, nil
}
