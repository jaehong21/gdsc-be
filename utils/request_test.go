package utils

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jaehong21/ga-be/config"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestJsonResp(t *testing.T) {
	// create a mock http.ResponseWriter
	w := httptest.NewRecorder()

	// define test values
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	user := User{Name: "John", Email: "john@example.com"}
	statusCode := http.StatusOK

	// call the JsonResp function
	JsonResp(w, user, statusCode)

	// check if response code is correct
	if w.Code != statusCode {
		t.Errorf("Expected status code %v, but got %v", statusCode, w.Code)
	}

	// check if Content-Type header is set to application/json
	expectedContentType := "application/json"
	if w.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("Expected Content-Type header %v, but got %v", expectedContentType, w.Header().Get("Content-Type"))
	}

	// check if response body is correct
	expectedResponseBody := `{"name":"John","email":"john@example.com"}`
	if strings.Join(strings.Fields(w.Body.String()), " ") != strings.Join(strings.Fields(expectedResponseBody), " ") {
		t.Errorf("Expected response body %v, but got %v", expectedResponseBody, w.Body.String())
	}

	// statusCode = http.StatusInternalServerError

}

func TestD(t *testing.T) {
	_ = godotenv.Load("../.env")
	connStr := os.Getenv("TEST_POSTGRES_URL")
	db := config.InitDatabase(connStr)
	defer db.Close()
	// create a test handler function
	testHandler := func(w http.ResponseWriter, r *http.Request, db *sql.DB) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, world!")
	}

	// create a request and response recorder for the test
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	// create a new handler using D and call it with the test request and response recorder
	handler := D(db, testHandler)
	handler.ServeHTTP(w, req)

	// check the response status code and body
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}
	if w.Body.String() != "Hello, world!" {
		t.Errorf("expected response body %q but got %q", "Hello, world!", w.Body.String())
	}
}
func TestValidateRequest(t *testing.T) {
	server := httptest.NewServer(ValidateRequest(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test" {
			t.Errorf("Expected to request '/test', got: %s", r.URL.Path)
		}
		// Check if the context values were set correctly
		userID := r.Context().Value("user_id").(int)
		schoolID := r.Context().Value("school_id").(int)
		if userID != 1 {
			t.Errorf("Expected user_id to be 1, but got %v", userID)
		}
		if schoolID != 1 {
			t.Errorf("Expected school_id to be 1, but got %v", schoolID)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})))
	defer server.Close()

	client := &http.Client{}
	// Create a mock req with an authorization header
	req, _ := http.NewRequest(http.MethodGet, server.URL+"/test", nil)

	resp, _ := client.Do(req)
	if resp.StatusCode != 401 {
		t.Errorf("Expected status code 401, but got %v", resp.StatusCode)
	}

	testToken, _ := GenerateJwt(1, ROLE_STUDENT)
	req.Header.Set("Authorization", "Bearer "+testToken)

	resp, _ = client.Do(req)
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %v", resp.StatusCode)
	}
}
