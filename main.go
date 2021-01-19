package main

import (
	"toh-echo/domain"
	"toh-echo/infrastructure"

	"github.com/labstack/echo/v4"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
	dsn = "host=localhost user=toh dbname=toh password=toh sslmode=disable"
)

func main() {
	dbInit()
	infrastructure.Init()
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}

func dbInit() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Migrator().CreateTable(domain.User{})
}
