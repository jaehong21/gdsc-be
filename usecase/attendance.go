package usecase

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jaehong21/ga-be/api/dto"
	"github.com/jaehong21/ga-be/entity"
	"github.com/jaehong21/ga-be/storage"
	"github.com/jaehong21/ga-be/utils"
)

func CreateAttendance(w http.ResponseWriter, db *sql.DB, studentID int, body dto.CreateAttendanceDto) {
	lecture, err := storage.FindOneLecture(db, body.LectureID)
	if err != nil {
		utils.JsonResp(w, errors.New("lecture not exist"), http.StatusBadRequest)
		return
	}

	_, err = storage.FindOneLectureStudent(db, studentID, lecture.ID)
	if err != nil {
		utils.JsonResp(w, errors.New("not attending this lecture"), http.StatusBadRequest)
		return
	}

	building, err := storage.FindOneBuilding(db, lecture.BuildingID)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}
	if err = building.VerifyRequestIP(body.RequestIP); err != nil {
		utils.JsonResp(w, errors.New("invalid request ip"), http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	attendance, err := storage.CreateAttendance(tx, studentID, body.LectureID, entity.STATUS_OK, body.RequestIP, body.RequestLocation)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	tx.Commit()
	utils.JsonResp(w, attendance, http.StatusCreated)
}

func UpdateAttendance(w http.ResponseWriter, db *sql.DB, attendanceID int, professorID int, body dto.UpdateAttendanceDto) {
	attendance, err := storage.FindOneAttendance(db, attendanceID)
	if err != nil {
		utils.JsonResp(w, errors.New("attendance not exist"), http.StatusBadRequest)
		return
	}

	_, err = validateLecture(db, professorID, attendance.LectureID)
	if err != nil {
		utils.JsonResp(w, err, http.StatusForbidden)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	attendance, err = storage.UpdateAttendance(tx, attendance.ID, body.Status, "PROFESSOR")
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	tx.Commit()
	utils.JsonResp(w, attendance, http.StatusCreated)
}

func FindProfessorAttendance(w http.ResponseWriter, db *sql.DB, professorID int, lectureID string) {
	_, err := validateLecture(db, professorID, lectureID)
	if err != nil {
		utils.JsonResp(w, err, http.StatusForbidden)
		return
	}

	resp := []map[string]interface{}{}
	studentIDs, err := storage.FindDistinctStudent(db, lectureID)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	for _, studentID := range studentIDs {
		student, err := storage.FindOneStudent(db, studentID)
		if err != nil {
			utils.JsonResp(w, err, http.StatusInternalServerError)
			return
		}

		attendances, err := storage.FindAttendanceByStudentAndLecture(db, studentID, lectureID)
		if err != nil {
			utils.JsonResp(w, err, http.StatusInternalServerError)
			return
		}
		item := map[string]interface{}{
			"student_id":  studentID,
			"name":        student.Name,
			"email":       student.Email,
			"major":       student.Major,
			"attendances": attendances,
		}
		resp = append(resp, item)
	}

	utils.JsonResp(w, resp, http.StatusOK)
}

func ValidateRequestIP(w http.ResponseWriter, db *sql.DB, lectureID string, ipAddress string) {
	lecture, err := storage.FindOneLecture(db, lectureID)
	if err != nil {
		utils.JsonResp(w, errors.New("no lecture exist"), http.StatusBadRequest)
		return
	}
	building, err := storage.FindOneBuilding(db, lecture.BuildingID)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	if err = building.VerifyRequestIP(ipAddress); err != nil {
		utils.JsonResp(w, false, http.StatusOK)
	} else {
		utils.JsonResp(w, true, http.StatusOK)
	}
}
