package usecase

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jaehong21/ga-be/api/dto"
	"github.com/jaehong21/ga-be/storage"
	"github.com/jaehong21/ga-be/utils"
)

func CreateProfessor(w http.ResponseWriter, db *sql.DB, body dto.CreateProfessorDto) {
	_, err := storage.FindOneProfessor(db, body.ID)
	if err == nil {
		utils.JsonResp(w, errors.New("professor already exist"), http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	hashPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	professor, err := storage.CreateProfessor(tx, body.ID, hashPassword, body.Name, body.Major, body.Phone, body.Email)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	// Return the saved user as JSON
	token, err := utils.GenerateJwt(professor.ID, utils.ROLE_PROFESSOR)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"access_token": token,
	}

	tx.Commit()
	utils.JsonResp(w, resp, http.StatusCreated)
}

func LoginProfessor(w http.ResponseWriter, db *sql.DB, body dto.LoginProfessorDto) {
	professor, err := storage.FindOneProfessor(db, body.ID)
	if err != nil {
		utils.JsonResp(w, errors.New("no professor found"), http.StatusBadRequest)
		return
	}

	if !professor.CompareHashAndPassword(body.Password) {
		utils.JsonResp(w, errors.New("invalid password"), http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateJwt(professor.ID, utils.ROLE_PROFESSOR)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"access_token": token,
	}

	utils.JsonResp(w, resp, http.StatusOK)
}
