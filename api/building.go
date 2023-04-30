package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jaehong21/ga-be/storage"
	"github.com/jaehong21/ga-be/utils"
)

func FindOneBuilding(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(vars)

	building, err := storage.FindOneBuilding(db, id)
	if err != nil {
		utils.JsonResp(w, errors.New("no building found"), http.StatusBadRequest)
		return
	}

	building.NetworkAddress = ""
	building.SubnetMask = ""

	utils.JsonResp(w, building, http.StatusOK)
}
