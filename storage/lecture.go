package storage

import (
	"database/sql"

	"github.com/jaehong21/ga-be/entity"
)

func CreateLecture(tx *sql.Tx, id string, name string, professorID int, lectureStartTime string, buildingID int, attendanceValidTime int) (entity.Lecture, error) {
	var lecture entity.Lecture
	row := tx.QueryRow("INSERT INTO \"lecture\" (id, name, professor_id, lecture_start_time, building_id, attendance_valid_time) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *", id, name, professorID, lectureStartTime, buildingID, attendanceValidTime)
	err := row.Scan(&lecture.ID, &lecture.Name, &lecture.ProfessorID, &lecture.LectureStartTime, &lecture.BuildingID, &lecture.AttendanceValidTime)
	if err != nil {
		return lecture, err
	}

	return lecture, nil
}

func FindLectureByStudent(db *sql.DB, studentID int) ([]entity.Lecture, error) {
	var lectures []entity.Lecture
	var lectureIDs []string

	rows, err := db.Query("SELECT lecture_id FROM \"lecture_student_rel\" WHERE student_id=$1", studentID)
	if err != nil {
		return lectures, err
	}

	for rows.Next() {
		var lectureID string
		err := rows.Scan(&lectureID)
		if err != nil {
			return lectures, err
		}
		lectureIDs = append(lectureIDs, lectureID)
	}

	for _, lectureID := range lectureIDs {
		var lecture entity.Lecture
		row := db.QueryRow("SELECT id, name, professor_id, lecture_start_time, building_id, attendance_valid_time FROM \"lecture\" WHERE id=$1", lectureID)
		err := row.Scan(&lecture.ID, &lecture.Name, &lecture.ProfessorID, &lecture.LectureStartTime, &lecture.BuildingID, &lecture.AttendanceValidTime)
		if err != nil {
			return lectures, err
		}
		lectures = append(lectures, lecture)
	}

	return lectures, nil
}

func FindLectureByProfessor(db *sql.DB, professorID int) ([]entity.Lecture, error) {
	var lectures []entity.Lecture

	rows, err := db.Query("SELECT id, name, professor_id, lecture_start_time, building_id, attendance_valid_time FROM \"lecture\" WHERE professor_id=$1", professorID)
	if err != nil {
		return lectures, err
	}

	for rows.Next() {
		var lecture entity.Lecture
		err := rows.Scan(&lecture.ID, &lecture.Name, &lecture.ProfessorID, &lecture.LectureStartTime, &lecture.BuildingID, &lecture.AttendanceValidTime)
		if err != nil {
			return lectures, err
		}
		lectures = append(lectures, lecture)
	}

	return lectures, nil
}

func FindOneLecture(db *sql.DB, id string) (entity.Lecture, error) {
	var lecture entity.Lecture
	row := db.QueryRow("SELECT id, name, professor_id, lecture_start_time, building_id, attendance_valid_time FROM \"lecture\" WHERE id=$1", id)
	err := row.Scan(&lecture.ID, &lecture.Name, &lecture.ProfessorID, &lecture.LectureStartTime, &lecture.BuildingID, &lecture.AttendanceValidTime)
	if err != nil {
		return lecture, err
	}

	return lecture, nil
}

func FindOneLectureStudent(db *sql.DB, studentID int, lectureID string) (entity.LectureStudent, error) {
	var lectureStudent entity.LectureStudent
	row := db.QueryRow("SELECT student_id, lecture_id FROM \"lecture_student_rel\" WHERE student_id=$1 AND lecture_id=$2", studentID, lectureID)
	err := row.Scan(&lectureStudent.StudentID, &lectureStudent.LectureID)
	if err != nil {
		return lectureStudent, err
	}

	return lectureStudent, nil
}

func FindLectureStudent(db *sql.DB, lectureID string) ([]entity.LectureStudent, error) {
	var lectureStudents []entity.LectureStudent

	rows, err := db.Query("SELECT student_id, lecture_id FROM \"lecture_student_rel\" WHERE lecture_id=$1", lectureID)
	if err != nil {
		return lectureStudents, err
	}

	for rows.Next() {
		var lectureStudent entity.LectureStudent
		err := rows.Scan(&lectureStudent.StudentID, &lectureStudent.LectureID)
		if err != nil {
			return lectureStudents, err
		}
		lectureStudents = append(lectureStudents, lectureStudent)
	}

	return lectureStudents, nil
}

func CreateLectureStudent(tx *sql.Tx, studentID int, lectureID string) (entity.LectureStudent, error) {
	var lectureStudent entity.LectureStudent
	row := tx.QueryRow("INSERT INTO \"lecture_student_rel\" (student_id, lecture_id) VALUES ($1, $2) RETURNING *", studentID, lectureID)
	err := row.Scan(&lectureStudent.StudentID, &lectureStudent.LectureID)
	if err != nil {
		return lectureStudent, err
	}

	return lectureStudent, nil
}

func DeleteLectureStudent(tx *sql.Tx, studentID int, lectureID string) (entity.LectureStudent, error) {
	var lectureStudent entity.LectureStudent
	row := tx.QueryRow("DELETE from \"lecture_student_rel\" WHERE student_id=$1 AND lecture_id=$2 RETURNING *", studentID, lectureID)
	err := row.Scan(&lectureStudent.StudentID, &lectureStudent.LectureID)
	if err != nil {
		return lectureStudent, err
	}

	return lectureStudent, nil
}
