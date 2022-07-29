package services

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/marvinarlt/goshort/internal/consts"
)

var DB *gorm.DB

func DbConnect() {
	var err error

	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	name := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True", user, pass, host, name)

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if nil != err {
		fmt.Println(consts.PREFIX, "Connecting to database failed")
		log.Fatalln(err.Error())
	}
}
