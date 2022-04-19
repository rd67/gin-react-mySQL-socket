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

type IAppConfig struct {
	Port int
	Name string
}

type IDatabaseConfig struct {
	Host     string
	Database string
	User     string
	Password string
	Port     string
}

type IJwtConfig struct {
	SecretKey    string
	HourLifespan int
}

type IConfig struct {
	App      IAppConfig
	Database IDatabaseConfig
	Jwt      IJwtConfig
}

var Config IConfig

func init() {

	Port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	AppConfig := IAppConfig{
		Port: Port,
		Name: os.Getenv("APP_NAME"),
	}

	DatabaseConfig := IDatabaseConfig{
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

	JwtConfig := IJwtConfig{
		SecretKey:    os.Getenv("JWT_SECRET_KEY"),
		HourLifespan: JwtHourLifespan,
	}

	Config = IConfig{
		App:      AppConfig,
		Database: DatabaseConfig,
		Jwt:      JwtConfig,
	}
}
