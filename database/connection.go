package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env파일을 불러오는중 오류가 발생하였습니다!")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBIP"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("데이터베이스 연결중 오류가 발생하였습니다!")
	}

	return db
}
