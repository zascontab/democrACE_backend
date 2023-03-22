package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sonderkevin/governance/infrastructure/log"
)

var (
	PORT                                  string
	DB_HOST                               string
	DB_PORT                               string
	DB_USER                               string
	DB_PASSWORD                           string
	DB_NAME                               string
	SECRET                                string
	ADMIN                                 string
	ADMIN_EMAIL                           string
	ADMIN_PASSWORD                        string
	TEMP_TOKEN_EXP                        string
	REFRESH_TOKEN_EXP                     string
	EMAIL_SENDER                          string
	EMAIL_SENDER_PASSWORD                 string
	EMAIL_HOST                            string
	EMAIL_PORT                            string
	MAX_SESIONES                          int
	REMOVE_EXPIRED_SESSIONS_TIME_INTERVAL string
)

var (
	ADMIN_DEFAULT_PERMISOS       []string
	FUNCIONARIO_DEFAULT_PERMISOS []string
	INVITADO_DEFAULT_PERMISOS    []string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	PORT = getEnv("PORT")
	DB_HOST = getEnv("DB_HOST")
	DB_PORT = getEnv("DB_PORT")
	DB_USER = getEnv("DB_USER")
	DB_PASSWORD = getEnv("DB_PASSWORD")
	DB_NAME = getEnv("DB_NAME")
	SECRET = getEnv("SECRET")
	ADMIN = getEnv("ADMIN")
	ADMIN_EMAIL = getEnv("ADMIN_EMAIL")
	ADMIN_PASSWORD = getEnv("ADMIN_PASSWORD")
	TEMP_TOKEN_EXP = getEnv("TEMP_TOKEN_EXP")
	REFRESH_TOKEN_EXP = getEnv("REFRESH_TOKEN_EXP")
	EMAIL_SENDER = getEnv("EMAIL_SENDER")
	EMAIL_SENDER_PASSWORD = getEnv("EMAIL_SENDER_PASSWORD")
	EMAIL_HOST = getEnv("EMAIL_HOST")
	EMAIL_PORT = getEnv("EMAIL_PORT")
	MAX_SESIONES, err = strconv.Atoi(getEnv("MAX_SESIONES"))
	if err != nil {
		log.Fatal("variable MAX_SESIONES debe ser de tipo int")
	}
	REMOVE_EXPIRED_SESSIONS_TIME_INTERVAL = getEnv("REMOVE_EXPIRED_SESSIONS_TIME_INTERVAL")

	ADMIN_DEFAULT_PERMISOS = []string{"all"}
	FUNCIONARIO_DEFAULT_PERMISOS = []string{"me"}
	INVITADO_DEFAULT_PERMISOS = []string{"me"}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal(fmt.Sprintf("%s environment variable not set", key))
	}
	return value
}
