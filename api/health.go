package api

import (
	"net/http"

	"github.com/jaehong21/ga-be/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}

func GetPublicKey(w http.ResponseWriter, r *http.Request) {
	publicKey, err := utils.GetPublicKey()
	if err != nil {
		utils.JsonResp(w, err, http.StatusInternalServerError)
		return
	}

	utils.JsonResp(w, publicKey, http.StatusOK)
}
