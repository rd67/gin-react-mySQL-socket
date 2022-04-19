package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type AppConfigStruct struct {
	Port int
	Name string
}

type DatabaseConfigStruct struct {
	Host     string
	Database string
	User     string
	Password string
	Port     string
}

type JwtConfigStruct struct {
	SecretKey    string
	HourLifespan int
}

type ConfigStruct struct {
	App      AppConfigStruct
	Database DatabaseConfigStruct
	Jwt      JwtConfigStruct
}

var Config ConfigStruct

func init() {

	Port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	AppConfig := AppConfigStruct{
		Port: Port,
		Name: os.Getenv("APP_NAME"),
	}

	DatabaseConfig := DatabaseConfigStruct{
		Host:     os.Getenv("DATABASE_HOST"),
		Database: os.Getenv("DATABASE_NAME"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Port:     os.Getenv("DATABASE_PORT"),
	}

	JwtHourLifespan, err := strconv.Atoi(os.Getenv("JWT_HOUR_LIFESPAN"))
	if err != nil {
		log.Fatal(err.Error())
	}

	JwtConfig := JwtConfigStruct{
		SecretKey:    os.Getenv("JWT_SECRET_KEY"),
		HourLifespan: JwtHourLifespan,
	}

	Config = ConfigStruct{
		App:      AppConfig,
		Database: DatabaseConfig,
		Jwt:      JwtConfig,
	}
}
