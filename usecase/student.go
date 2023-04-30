package usecase

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jaehong21/ga-be/api/dto"
	"github.com/jaehong21/ga-be/storage"
	"github.com/jaehong21/ga-be/utils"
)

func CreateStudent(w http.ResponseWriter, db *sql.DB, body dto.CreateStudentDto) {
	_, err := storage.FindOneStudent(db, body.ID)
	if err == nil {
		utils.JsonResp(w, errors.New("user already exist"), http.StatusBadRequest)
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

	student, err := storage.CreateStudent(tx, body.ID, hashPassword, body.Name, body.Major, body.Phone, body.Email, body.DeviceID)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	// Return the saved user as JSON
	token, err := utils.GenerateJwt(student.ID, utils.ROLE_STUDENT)
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

func LoginStudent(w http.ResponseWriter, db *sql.DB, body dto.LoginStudentDto) {
	student, err := storage.FindOneStudent(db, body.ID)
	if err != nil {
		utils.JsonResp(w, errors.New("no user found"), http.StatusBadRequest)
		return
	}

	if !student.CompareHashAndPassword(body.Password) {
		utils.JsonResp(w, errors.New("invalid password"), http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateJwt(student.ID, utils.ROLE_STUDENT)
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"access_token": token,
	}

	utils.JsonResp(w, resp, http.StatusOK)
}
