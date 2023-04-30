package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	envVars := []string{"POSTGRES_URL", "TEST_POSTGRES_URL", "LISTEN_ADDR", "JWT_SECRET_KEY"}
	for _, v := range envVars {
		if os.Getenv(v) == "" {
			log.Fatalf("environment variable %s is not set", v)
		}
	}
}

func InitDatabase(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
