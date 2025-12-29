package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"hotlog.org/db"
	"hotlog.org/handlers"
	"hotlog.org/repository"
	"hotlog.org/service"
)

func main() {
	gormDB, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.ConfigurePool(gormDB)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = sqlDB.Close()
	}()

	authRepo := repository.NewAuthRepository(gormDB)
	authService := service.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	router := gin.Default()
	authHandler.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
