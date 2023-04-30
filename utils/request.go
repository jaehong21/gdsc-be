package utils

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jaehong21/ga-be/entity"
)

var validate *validator.Validate = validator.New()

func ParseJsonBody(r *http.Request, body any) error {
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		return err
	}

	validate.RegisterValidation("phone", validatePhone)
	validate.RegisterValidation("lecturetime", validateLectureTimeFormat)
	validate.RegisterValidation("attendancestatus", validateAttendnaceStatus)
	err = validate.Struct(body)
	if err != nil {
		return err
	}

	return nil
}

func validatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	pattern := `^010\d{8}$`
	match, _ := regexp.MatchString(pattern, phone)
	return match
}
func validateLectureTimeFormat(fl validator.FieldLevel) bool {
	lectureTime := fl.Field().String()
	pattern := `^\d{2}:\d{2}$`
	match, _ := regexp.MatchString(pattern, lectureTime)
	return match
}
func validateAttendnaceStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	switch status {
	case entity.STATUS_OK, entity.STATUS_LATE, entity.STATUS_ABSENT, entity.STATUS_EXCUSED, entity.STATUS_UNKNOWN, entity.STATUS_ETC:
		return true
	default:
		return false
	}
}

func JsonResp(w http.ResponseWriter, value any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	switch v := value.(type) {
	case error:
		resp := map[string]interface{}{
			"msg": v.Error(),
		}
		json.NewEncoder(w).Encode(resp)
	case string:
		resp := map[string]interface{}{
			"msg": v,
		}
		json.NewEncoder(w).Encode(resp)
	default:
		json.NewEncoder(w).Encode(value)
	}
}

func D(db *sql.DB, fn func(http.ResponseWriter, *http.Request, *sql.DB)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	})
}

func ValidateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			JsonResp(w, errors.New("authorization header missing"), http.StatusUnauthorized)
			return
		}

		// Split the "Bearer " prefix from the token string
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			JsonResp(w, errors.New("invalid Authorization header format"), http.StatusUnauthorized)
			return
		}
		tokenString = tokenParts[1]

		claims, err := ValidateJwt(tokenString)
		if err != nil {
			JsonResp(w, err, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "sub", claims.UserID)
		ctx = context.WithValue(ctx, "role", claims.Role)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
