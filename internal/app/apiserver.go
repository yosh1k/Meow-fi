package app

import (
	"Meow-fi/internal/config"
	"Meow-fi/internal/models"
	"Meow-fi/internal/transport"
	"fmt"

	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	state bool
}

func (server *Server) Start() {
	dbinit()
	transport.Init()
	e := echo.New()
	e.Logger.Fatal(e.Start(config.ServerPort))
}

func dbinit() {

	db, err := gorm.Open(postgres.Open(config.DatabaseUrl), &gorm.Config{})

	err = db.Migrator().CreateTable(models.User{})
	if err != nil {
		fmt.Print("User already exists")
	}
	err = db.Migrator().CreateTable(models.Notice{})
	if err != nil {
		fmt.Print("Notice already exists")
	}
	err = db.Migrator().CreateTable(models.Deal{})
	if err != nil {
		fmt.Print("Deal already exists")
	}
}
