package storage

import (
	"database/sql"

	"github.com/jaehong21/ga-be/entity"
)

func CreateAttendance(tx *sql.Tx, studentID int, lectureID string, status string, requestIP string, requestLocation string) (entity.Attendance, error) {
	var attendance entity.Attendance
	row := tx.QueryRow("INSERT INTO \"attendance\" (student_id, lecture_id, status, request_ip, request_location) VALUES ($1, $2, $3, $4, $5) RETURNING *", studentID, lectureID, status, requestIP, requestLocation)
	err := row.Scan(&attendance.ID, &attendance.StudentID, &attendance.LectureID, &attendance.Status, &attendance.RequestIP, &attendance.RequestLocation, &attendance.Validator, &attendance.CreatedAt)
	if err != nil {
		return attendance, err
	}

	return attendance, nil
}

func UpdateAttendance(tx *sql.Tx, id int, status string, validator string) (entity.Attendance, error) {
	var attendance entity.Attendance
	row := tx.QueryRow("UPDATE \"attendance\" SET status=$1, validator=$2 WHERE id=$3 RETURNING *", status, validator, id)
	err := row.Scan(&attendance.ID, &attendance.StudentID, &attendance.LectureID, &attendance.Status, &attendance.RequestIP, &attendance.RequestLocation, &attendance.Validator, &attendance.CreatedAt)
	if err != nil {
		return attendance, err
	}

	return attendance, nil
}

func FindOneAttendance(db *sql.DB, id int) (entity.Attendance, error) {
	var attendance entity.Attendance
	row := db.QueryRow("SELECT id, student_id, lecture_id, status, request_ip, request_location, validator, created_at FROM \"attendance\" WHERE id=$1", id)
	err := row.Scan(&attendance.ID, &attendance.StudentID, &attendance.LectureID, &attendance.Status, &attendance.RequestIP, &attendance.RequestLocation, &attendance.Validator, &attendance.CreatedAt)
	if err != nil {
		return attendance, err
	}

	return attendance, nil
}

func FindAttendanceByStudentAndLecture(db *sql.DB, studentID int, lectureID string) ([]entity.Attendance, error) {
	var attendances []entity.Attendance
	rows, err := db.Query("SELECT id, student_id, lecture_id, status, request_ip, request_location, validator, created_at FROM \"attendance\" WHERE student_id=$1 AND lecture_id=$2 ORDER BY created_at DESC", studentID, lectureID)
	if err != nil {
		return attendances, err
	}
	for rows.Next() {
		var attendance entity.Attendance
		err := rows.Scan(&attendance.ID, &attendance.StudentID, &attendance.LectureID, &attendance.Status, &attendance.RequestIP, &attendance.RequestLocation, &attendance.Validator, &attendance.CreatedAt)
		if err != nil {
			return attendances, err
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}

func FindAttendanceByLecture(db *sql.DB, lectureID string) ([]entity.Attendance, error) {
	var attendances []entity.Attendance
	rows, err := db.Query("SELECT id, student_id, lecture_id, status, request_ip, request_location, validator, created_at FROM \"attendance\" WHERE lecture_id=$1 ORDER BY created_at", lectureID)
	if err != nil {
		return attendances, err
	}
	for rows.Next() {
		var attendance entity.Attendance
		err := rows.Scan(&attendance.ID, &attendance.StudentID, &attendance.LectureID, &attendance.Status, &attendance.RequestIP, &attendance.RequestLocation, &attendance.Validator, &attendance.CreatedAt)
		if err != nil {
			return attendances, err
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}

func CountAttendanceByStudent(db *sql.DB, studentID int, lectureID string) (int, int, int, error) {
	ok, late, absent := 0, 0, 0
	stmt := "SELECT COUNT(*) FROM \"attendance\" WHERE student_id=$1 AND lecture_id=$2 AND status=$3"

	row := db.QueryRow(stmt, studentID, lectureID, entity.STATUS_OK)
	err := row.Scan(&ok)
	if err != nil {
		return ok, late, absent, err
	}

	row = db.QueryRow(stmt, studentID, lectureID, entity.STATUS_LATE)
	err = row.Scan(&late)
	if err != nil {
		return ok, late, absent, err
	}

	row = db.QueryRow(stmt, studentID, lectureID, entity.STATUS_ABSENT)
	err = row.Scan(&absent)
	if err != nil {
		return ok, late, absent, err
	}

	return ok, late, absent, nil
}

func FindDistinctStudent(db *sql.DB, lectureID string) ([]int, error) {
	var students []int
	rows, err := db.Query("SELECT DISTINCT student_id FROM \"attendance\" WHERE lecture_id=$1", lectureID)
	if err != nil {
		return students, err
	}
	for rows.Next() {
		var student int
		err := rows.Scan(&student)
		if err != nil {
			return students, err
		}
		students = append(students, student)
	}

	return students, nil
}

func FindDistinctTime(db *sql.DB, lectureID string) ([]string, error) {
	var times []string
	rows, err := db.Query("SELECT DISTINCT date_trunc('day', created_at) FROM \"attendance\" WHERE lecture_id=$1", lectureID)
	if err != nil {
		return times, err
	}
	for rows.Next() {
		var time string
		err := rows.Scan(&time)
		if err != nil {
			return times, err
		}
		times = append(times, time[:10])
	}

	return times, nil
}
