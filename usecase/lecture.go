package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/jaehong21/ga-be/api/dto"
	"github.com/jaehong21/ga-be/entity"
	"github.com/jaehong21/ga-be/storage"
	"github.com/jaehong21/ga-be/utils"
)

func FindLectureStudent(w http.ResponseWriter, db *sql.DB, lectureID string) {
	lectureStudents, err := storage.FindLectureStudent(db, lectureID)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	var students []entity.Student
	for _, lectureStudent := range lectureStudents {
		student, err := storage.FindOneStudent(db, lectureStudent.StudentID)
		if err != nil {
			utils.JsonResp(w, err, http.StatusInternalServerError)
			return
		}
		student.Password = ""
		students = append(students, student)
	}

	utils.JsonResp(w, students, http.StatusOK)
}

func CreateLectureStudent(w http.ResponseWriter, db *sql.DB, professorID int, body dto.CreateLectureStudentDto) {
	lecture, err := storage.FindOneLecture(db, body.LectureID)
	if err != nil {
		utils.JsonResp(w, errors.New("lecture not exist"), http.StatusBadRequest)
		return
	}

	if lecture.ProfessorID != professorID {
		utils.JsonResp(w, errors.New("not instructor of this lecture"), http.StatusForbidden)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	count := 0
	for _, studentID := range body.StudentIDs {
		_, err = storage.FindOneStudent(db, studentID)
		if err != nil {
			utils.JsonResp(w, errors.New("student not exist"), http.StatusBadRequest)
			return
		}
		_, err = storage.CreateLectureStudent(tx, studentID, body.LectureID)
		if err != nil {
			utils.JsonResp(w, err, http.StatusInternalServerError)
			return
		}
		count++
	}

	tx.Commit()
	msg := fmt.Sprintf("success, %d students added", count)
	utils.JsonResp(w, msg, http.StatusCreated)
}

func DeleteLectureStudent(w http.ResponseWriter, db *sql.DB, professorID int, body dto.DeleteLectureStudent) {
	_, err := storage.FindOneLectureStudent(db, body.StudentID, body.LectureID)
	if err != nil {
		utils.JsonResp(w, errors.New("student from this lecture not exist"), http.StatusBadRequest)
		return
	}

	lecture, err := storage.FindOneLecture(db, body.LectureID)
	if err != nil {
		utils.JsonResp(w, errors.New("lecture not exist"), http.StatusBadRequest)
		return
	}

	if lecture.ProfessorID != professorID {
		utils.JsonResp(w, errors.New("not instructor of this lecture"), http.StatusForbidden)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	_, err = storage.DeleteLectureStudent(tx, body.StudentID, body.LectureID)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	tx.Commit()
	utils.JsonResp(w, "success", http.StatusOK)
}

func CreateLecture(w http.ResponseWriter, db *sql.DB, professorID int, body dto.CreateLectureDto) {
	_, err := storage.FindOneLecture(db, body.ID)
	if err == nil {
		utils.JsonResp(w, errors.New("lecture already exists"), http.StatusBadRequest)
		return
	}

	_, err = storage.FindOneBuilding(db, body.BuildingID)
	if err != nil {
		utils.JsonResp(w, errors.New("building not exist"), http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	lecture, err := storage.CreateLecture(tx, body.ID, body.Name, professorID, body.LectureStartTime, body.BuildingID, body.AttendanceValidTime)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	tx.Commit()
	utils.JsonResp(w, lecture, http.StatusCreated)
}

func validateLecture(db *sql.DB, professorID int, lectureID string) (entity.Lecture, error) {
	var lecture entity.Lecture
	lecture, err := storage.FindOneLecture(db, lectureID)
	if err != nil {
		return lecture, err
	}
	if lecture.ProfessorID != professorID {
		return lecture, errors.New("not instructor of this lecture")
	}

	return lecture, nil
}
