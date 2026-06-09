package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Jwt struct {
	Secret    string
	ExpiredAt int
}

type Config struct {
	Port     int
	Database Database
	Jwt      Jwt
}

var Data Config

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error when loading .env file: %v", err)
	}
	fmt.Println("Environment data loaded...")
}

func Load() {
	appPort, _ := strconv.Atoi(GetEnv("PORT", "8000"))
	dbPort, _ := strconv.Atoi(GetEnv("DB_PORT", "5432"))
	jwtExpiration, _ := strconv.Atoi(GetEnv("JWT_EXPIRATION", "604800"))

	Data = Config{
		Port: appPort,
		Database: Database{
			Host:     GetEnv("DB_HOST", "postgres"),
			Port:     dbPort,
			User:     GetEnv("DB_USER", "postgres"),
			Password: GetEnv("DB_PASSWORD", "postgres"),
			DBName:   GetEnv("DB_NAME", "coonfiber"),
		},
		Jwt: Jwt{
			Secret:    GetEnv("JWT_SECRET", "themostsecret"),
			ExpiredAt: jwtExpiration, // Default 1 week
		},
	}
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
