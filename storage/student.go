package storage

import (
	"database/sql"

	"github.com/jaehong21/ga-be/entity"
)

func CreateStudent(tx *sql.Tx, id int, password string, name string, major string, phone string, email string, deviceID string) (entity.Student, error) {
	var student entity.Student

	row := tx.QueryRow("INSERT INTO \"student\" (id, password, name, major, phone, email, device_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *", id, password, name, major, phone, email, deviceID)
	err := row.Scan(&student.ID, &student.Password, &student.Name, &student.Major, &student.Phone, &student.Email, &student.DeviceID)
	if err != nil {
		return student, err
	}

	return student, nil
}

func FindOneStudent(db *sql.DB, id int) (entity.Student, error) {
	var student entity.Student

	row := db.QueryRow("SELECT id, password, name, major, phone, email, device_id FROM \"student\" WHERE id=$1", id)
	err := row.Scan(&student.ID, &student.Password, &student.Name, &student.Major, &student.Phone, &student.Email, &student.DeviceID)
	if err != nil {
		return student, err
	}

	return student, nil
}
