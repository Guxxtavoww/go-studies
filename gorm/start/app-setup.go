package start

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DB_PORT                string
	DB_USER                string
	DATABASE_ROOT_PASSWORD string
	DATABASE_DATABASE_NAME string
}

const ENV_CONTEXT_KEY = "env-context-key"

func (env *Env) GetDbConnectionDsn() string {
	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.DB_USER,
		env.DATABASE_ROOT_PASSWORD,
		env.DATABASE_DATABASE_NAME,
		env.DB_PORT,
	)

	return dsn
}

func AppSetup() *Env {
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
