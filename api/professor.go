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

func FindProfessorInfo(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := r.Context().Value("sub").(int)

	professor, err := storage.FindOneProfessor(db, id)
	if err != nil {
		utils.JsonResp(w, errors.New("no professor found"), http.StatusBadRequest)
		return
	}

	utils.JsonResp(w, professor, http.StatusOK)
}

func FindOneProfessor(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(vars)

	professor, err := storage.FindOneProfessor(db, id)
	if err != nil {
		utils.JsonResp(w, errors.New("no professor found"), http.StatusBadRequest)
		return
	}
	// prevent hash_password from being sent to client
	professor.Password = ""

	utils.JsonResp(w, professor, http.StatusOK)
}

func CreateProfessor(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var body dto.CreateProfessorDto
	err := utils.ParseJsonBody(r, &body)
	if err != nil {
		utils.JsonResp(w, err, http.StatusBadRequest)
		return
	}

	usecase.CreateProfessor(w, db, body)
}

func LoginProfessor(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var body dto.LoginProfessorDto
	err := utils.ParseJsonBody(r, &body)
	if err != nil {
		utils.JsonResp(w, err, http.StatusBadRequest)
		return
	}

	usecase.LoginProfessor(w, db, body)
}
