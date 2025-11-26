package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB เป็นตัวแปร Global เพื่อให้ที่อื่นเรียกใช้ได้ (หรือจะส่งผ่าน Struct ก็ได้)
var DB *gorm.DB

func ConnectDb() {
	var err error

	// อ่านค่าจาก Environment Variable
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	// เชื่อมต่อ Database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected to database successfully")

	// (Optional) ทำ Auto Migrate เพื่อสร้างตารางอัตโนมัติ
	// DB.AutoMigrate(&models.User{})
}
