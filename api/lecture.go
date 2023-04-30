package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaehong21/ga-be/api/dto"
	"github.com/jaehong21/ga-be/entity"
	"github.com/jaehong21/ga-be/storage"
	"github.com/jaehong21/ga-be/usecase"
	"github.com/jaehong21/ga-be/utils"
)

func FindOneLecture(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := mux.Vars(r)["id"]

	lecture, err := storage.FindOneLecture(db, id)
	if err != nil {
		utils.JsonResp(w, errors.New("no lecture found"), http.StatusBadRequest)
		return
	}

	utils.JsonResp(w, lecture, http.StatusOK)
}

func FindLecture(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := r.Context().Value("sub").(int)
	role := r.Context().Value("role").(string)

	var lectures []entity.Lecture

	switch role {
	case utils.ROLE_STUDENT:
		result, err := storage.FindLectureByStudent(db, id)
		if err != nil {
			utils.JsonResp(w, err, http.StatusInternalServerError)
			return
		}
		lectures = result
	case utils.ROLE_PROFESSOR:
		result, err := storage.FindLectureByProfessor(db, id)
		if err != nil {
			utils.JsonResp(w, err, http.StatusInternalServerError)
			return
		}
		lectures = result
	default:
		utils.JsonResp(w, errors.New("invalid role"), http.StatusUnauthorized)
	}

	utils.JsonResp(w, lectures, http.StatusOK)
}

func CreateLecture(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := r.Context().Value("sub").(int)
	if r.Context().Value("role").(string) != utils.ROLE_PROFESSOR {
		utils.JsonResp(w, "only professor can create lecture", http.StatusForbidden)
		return
	}

	var body dto.CreateLectureDto
	err := utils.ParseJsonBody(r, &body)
	if err != nil {
		utils.JsonResp(w, err, http.StatusBadRequest)
		return
	}

	usecase.CreateLecture(w, db, id, body)
}

func FindLectureStudent(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	lectureID := mux.Vars(r)["id"]

	usecase.FindLectureStudent(w, db, lectureID)
}

func CreateLectureStudent(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := r.Context().Value("sub").(int)
	if r.Context().Value("role").(string) != utils.ROLE_PROFESSOR {
		utils.JsonResp(w, "only professor can add student", http.StatusForbidden)
		return
	}

	var body dto.CreateLectureStudentDto
	err := utils.ParseJsonBody(r, &body)
	if err != nil {
		utils.JsonResp(w, err, http.StatusBadRequest)
		return
	}

	usecase.CreateLectureStudent(w, db, id, body)
}

func DeleteLectureStudent(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := r.Context().Value("sub").(int)
	if r.Context().Value("role").(string) != utils.ROLE_PROFESSOR {
		utils.JsonResp(w, "only professor can delete student", http.StatusForbidden)
		return
	}

	var body dto.DeleteLectureStudent
	err := utils.ParseJsonBody(r, &body)
	if err != nil {
		utils.JsonResp(w, err, http.StatusBadRequest)
		return
	}

	usecase.DeleteLectureStudent(w, db, id, body)
}

func FindDistinctTime(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	lectureID := mux.Vars(r)["id"]

	times, err := storage.FindDistinctTime(db, lectureID)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	utils.JsonResp(w, times, http.StatusOK)
}
