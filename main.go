package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jaehong21/ga-be/api"
	"github.com/jaehong21/ga-be/config"
	. "github.com/jaehong21/ga-be/utils"
	"github.com/rs/cors"
)

func main() {
	config.LoadEnv()

	db := config.InitDatabase(os.Getenv("POSTGRES_URL"))
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/", api.HealthCheck).Methods(http.MethodGet)
	router := r.PathPrefix("/v1").Subrouter()

	keyRouter := router.PathPrefix("/key").Subrouter()
	keyRouter.HandleFunc("/public", api.GetPublicKey).Methods(http.MethodGet)

	studentRouter := router.PathPrefix("/student").Subrouter()
	studentRouter.Handle("", D(db, api.CreateStudent)).Methods(http.MethodPost)
	studentRouter.Handle("/login", D(db, api.LoginStudent)).Methods(http.MethodPost)
	studentRouter.Handle("/{id:[0-9]+}", D(db, api.FindOneUser)).Methods(http.MethodGet)
	studentRouter.Handle("/info", ValidateRequest(D(db, api.FindUserInfo))).Methods(http.MethodGet)

	professorRouter := router.PathPrefix("/professor").Subrouter()
	professorRouter.Handle("", D(db, api.CreateProfessor)).Methods(http.MethodPost)
	professorRouter.Handle("/login", D(db, api.LoginProfessor)).Methods(http.MethodPost)
	professorRouter.Handle("/{id:[0-9]+}", D(db, api.FindOneProfessor)).Methods(http.MethodGet)
	professorRouter.Handle("/info", ValidateRequest(D(db, api.FindProfessorInfo))).Methods(http.MethodGet)

	buildingRouter := router.PathPrefix("/building").Subrouter()
	buildingRouter.Handle("/{id:[0-9]+}", D(db, api.FindOneBuilding)).Methods(http.MethodGet)

	lectureRouter := router.PathPrefix("/lecture").Subrouter()
	lectureRouter.Handle("", ValidateRequest(D(db, api.CreateLecture))).Methods(http.MethodPost)
	lectureRouter.Handle("", ValidateRequest(D(db, api.FindLecture))).Methods(http.MethodGet)
	lectureRouter.Handle("/{id}", D(db, api.FindOneLecture)).Methods(http.MethodGet)
	lectureRouter.Handle("/{id}/student/list", ValidateRequest(D(db, api.FindLectureStudent))).Methods(http.MethodGet)
	lectureRouter.Handle("/student/add", ValidateRequest(D(db, api.CreateLectureStudent))).Methods(http.MethodPost)
	lectureRouter.Handle("/student/del", ValidateRequest(D(db, api.DeleteLectureStudent))).Methods(http.MethodDelete)
	lectureRouter.Handle("/timetable/{id}", ValidateRequest(D(db, api.FindDistinctTime))).Methods(http.MethodGet)

	attendanceRouter := router.PathPrefix("/attendance").Subrouter()
	attendanceRouter.Handle("", ValidateRequest(D(db, api.CreateAttendance))).Methods(http.MethodPost)
	attendanceRouter.Handle("/{id:[0-9]+}", ValidateRequest(D(db, api.UpdateAttendance))).Methods(http.MethodPatch)
	attendanceRouter.Handle("/student/lecture/{id}", ValidateRequest(D(db, api.FindStudentAttendance))).Methods(http.MethodGet)
	attendanceRouter.Handle("/professor/lecture/{id}", ValidateRequest(D(db, api.FindProfessorAttendance))).Methods(http.MethodGet)
	attendanceRouter.Handle("/student/lecture/{id}/count", ValidateRequest(D(db, api.CountStudentAttendance))).Methods(http.MethodGet)
	attendanceRouter.Handle("/lecture/{id}/validate/ip/{address}", ValidateRequest(D(db, api.ValidateRequestIP))).Methods(http.MethodGet)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://wa-khaki.vercel.app"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(r)

	server := &http.Server{
		Addr: os.Getenv("LISTEN_ADDR"),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler, // Pass our instance of gorilla/mux in.
	}

	log.Println("Server is running on", os.Getenv("LISTEN_ADDR"))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
