package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Env struct {
	DB_PORT                string
	DB_USER                string
	DATABASE_ROOT_PASSWORD string
	DATABASE_DATABASE_NAME string
}

const ENV_CONTEXT_KEY = "env-context-key"

func appSetup() *Env {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env")
	}

	env := &Env{
		DB_PORT:                os.Getenv("DB_PORT"),
		DB_USER:                os.Getenv("DB_USER"),
		DATABASE_ROOT_PASSWORD: os.Getenv("DATABASE_ROOT_PASSWORD"),
		DATABASE_DATABASE_NAME: os.Getenv("DATABASE_DATABASE_NAME"),
	}

	return env
}

func main() {
	envVariables := appSetup()
	appContext := context.Background()
	appContext = context.WithValue(appContext, ENV_CONTEXT_KEY, envVariables)

	databaseConnectionString := fmt.Sprintf(
		"host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		envVariables.DB_PORT,
		envVariables.DB_USER,
		envVariables.DATABASE_ROOT_PASSWORD,
		envVariables.DATABASE_DATABASE_NAME,
	)

	db, err := sql.Open("postgres", databaseConnectionString)

	if err != nil {
		panic("Could not connect to database")
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
