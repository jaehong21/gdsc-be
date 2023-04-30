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

func FindUserInfo(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := r.Context().Value("sub").(int)

	student, err := storage.FindOneStudent(db, id)
	if err != nil {
		utils.JsonResp(w, errors.New("no student found"), http.StatusBadRequest)
		return
	}

	utils.JsonResp(w, student, http.StatusOK)
}

func FindOneUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(vars)

	student, err := storage.FindOneStudent(db, id)
	if err != nil {
		utils.JsonResp(w, errors.New("no student found"), http.StatusBadRequest)
		return
	}
	// prevent hash_password from being sent to client
	student.Password = ""

	utils.JsonResp(w, student, http.StatusOK)
}

func CreateStudent(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var body dto.CreateStudentDto
	err := utils.ParseJsonBody(r, &body)
	if err != nil {
		utils.JsonResp(w, err, http.StatusBadRequest)
		return
	}

	usecase.CreateStudent(w, db, body)
}

func LoginStudent(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var body dto.LoginStudentDto
	err := utils.ParseJsonBody(r, &body)
	if err != nil {
		utils.JsonResp(w, err, http.StatusBadRequest)
		return
	}

	usecase.LoginStudent(w, db, body)
}
