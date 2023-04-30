package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jaehong21/ga-be/api/dto"
	"github.com/jaehong21/ga-be/storage"
	"github.com/jaehong21/ga-be/usecase"
	"github.com/jaehong21/ga-be/utils"
)

func CreateAttendance(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := r.Context().Value("sub").(int)

	var body dto.CreateAttendanceDto
	err := utils.ParseJsonBody(r, &body)
	if err != nil {
		utils.JsonResp(w, err, http.StatusBadRequest)
		return
	}

	usecase.CreateAttendance(w, db, id, body)
}

func UpdateAttendance(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(vars)

	professorID := r.Context().Value("sub").(int)
	if r.Context().Value("role").(string) != utils.ROLE_PROFESSOR {
		utils.JsonResp(w, errors.New("only professor can update attendance"), http.StatusForbidden)
		return
	}

	var body dto.UpdateAttendanceDto
	err := utils.ParseJsonBody(r, &body)
	if err != nil {
		utils.JsonResp(w, err, http.StatusBadRequest)
		return
	}

	usecase.UpdateAttendance(w, db, id, professorID, body)
}

func FindStudentAttendance(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	studentID := r.Context().Value("sub").(int)
	lectureID := mux.Vars(r)["id"]

	attendances, err := storage.FindAttendanceByStudentAndLecture(db, studentID, lectureID)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	utils.JsonResp(w, attendances, http.StatusOK)
}

func FindProfessorAttendance(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	lectureID := mux.Vars(r)["id"]

	professorID := r.Context().Value("sub").(int)
	if r.Context().Value("role").(string) != utils.ROLE_PROFESSOR {
		utils.JsonResp(w, errors.New("only professor can update attendance"), http.StatusForbidden)
		return
	}

	usecase.FindProfessorAttendance(w, db, professorID, lectureID)
}

func CountStudentAttendance(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	studentID := r.Context().Value("sub").(int)
	lectureID := mux.Vars(r)["id"]

	ok, late, absent, err := storage.CountAttendanceByStudent(db, studentID, lectureID)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"OK":     ok,
		"LATE":   late,
		"ABSENT": absent,
	}

	utils.JsonResp(w, resp, http.StatusOK)
}

func ValidateRequestIP(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	lectureID := mux.Vars(r)["id"]
	ipAddress := mux.Vars(r)["address"]

	usecase.ValidateRequestIP(w, db, lectureID, ipAddress)
}
