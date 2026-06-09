package database

import (
	"fmt"
	"log"

	"gocoon/core/config"
	"gocoon/core/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB Instance
var DB *gorm.DB

// Connect
func Connect() {
	var dsn string = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", config.Data.Database.Host, config.Data.Database.User, config.Data.Database.Password, config.Data.Database.DBName, config.Data.Database.Port)

	db_instance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	DB = db_instance

	fmt.Println("Database connected...")
}

// Migrate
func Migrate() {
	for _, model := range models.ModelList {
		DB.AutoMigrate(model)
	}

	fmt.Println("Database schema migrated...")
}
