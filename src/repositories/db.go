package repositories

import (
	"fmt"
	"log"
	"os"
	"petProj/src/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func NewDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("failed to load env file: %w", err)
	}
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host = localhost user=%s password=%s dbname=%s ", db_user, db_password, db_name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!=nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	err = db.AutoMigrate(&models.Employee{},&models.DailyAttendanceLog{})
	if err != nil {
		return nil, err
	}
	log.Println("Database Connected and tables migrated successfully")
	return db, nil
}