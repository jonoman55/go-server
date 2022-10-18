package main

import (
	"go-server/cfg"
	"go-server/db"
	"go-server/models"
)

func init() {
	cfg.LoadEnvVars()
	db.ConnectDB()
}

func main() {
	db.DB.AutoMigrate(&models.Post{})
}
