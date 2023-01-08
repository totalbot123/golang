package dao

import (
	"encoding/json"
	"fmt"
	"location-processing/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadUsers(users *[]models.LocationProcessing) {
	data, err := os.ReadFile("/location-processing/milanjecmenica/Documents/personal/gridu_courses/golang/location-processing/models/data.json")
	check(err)
	err = json.Unmarshal(data, &users)
	check(err)
}

func ConnectDatabase() {
	// db_hostname := os.Getenv("POSTGRES_HOST")
	// db_name := os.Getenv("POSTGRES_DB")
	// db_user := os.Getenv("POSTGRES_USER")
	// db_pass := os.Getenv("POSTGRES_PASSWORD")
	// db_port := os.Getenv("POSTGRES_PORT")
	db_hostname := "127.0.0.1"
	db_name := "location_processing"
	db_user := "postgres"
	db_pass := "mysecretpassword"
	db_port := "5432"

	dbURl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_user, db_pass, db_hostname, db_port, db_name)
	database, err := gorm.Open(postgres.Open(dbURl), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.LocationProcessing{})
	// var users []models.LocationProcessing
	// loadUsers(&users)
	// database.Create(&users)

	DB = database
}
