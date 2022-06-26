package main

import (
	"go_jwt"
	"go_jwt/database/config"
	"go_jwt/database/drop"
	"go_jwt/database/migrate"
	"go_jwt/database/seed"
	"go_jwt/helper"
	"go_jwt/pkg/handler"
	"go_jwt/pkg/repository"
	"go_jwt/pkg/service"
	"log"
)

func main() {
	db, err := config.DBConnect()
	helper.PanicIfErrorWithMessage("Error when connecting to database:", err)
	defer db.Close()

	drop.DropAllTable(db)
	migrate.Migrate(db)
	seed.Seed(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(go_jwt.Server)
	if err := srv.Run("8080", handlers.AllowOrigin(handlers.InitRoutes())); err != nil {
		log.Fatal("Error occured while running server: ", err.Error())
	}
}
