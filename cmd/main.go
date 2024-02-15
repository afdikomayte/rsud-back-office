package main

import (
	"log"

	"github.com/afdikomayte/rsud-back-office/external/database"
	"github.com/afdikomayte/rsud-back-office/internal/config"
)

func main() {
	filename := "cmd/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectMysql(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}
}
