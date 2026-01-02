package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hotlog.org/db"
	"hotlog.org/handlers"
	"hotlog.org/middleware"
	"hotlog.org/repository"
	"hotlog.org/service"
)

func main() {
	godotenv.Load()

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

	healthHandler := handlers.NewHealthHandler(sqlDB)

	projectRepo := repository.NewProjectRepository(gormDB)
	projectService := service.NewProjectService(projectRepo)
	projectHandler := handlers.NewProjectHandler(projectService)

	secret := os.Getenv("INTERNAL_API_SECRET")

	if secret == "" {
		log.Fatal("INTERNAL_API_SECRET is required")
	}

	maxSkewSeconds := int64(60)

	if rawSkew := os.Getenv("INTERNAL_API_MAX_SKEW_SECONDS"); rawSkew != "" {
		parsed, err := strconv.ParseInt(rawSkew, 10, 64)

		if err != nil {
			log.Fatal(err)
		}

		maxSkewSeconds = parsed
	}

	router := gin.Default()
	healthHandler.RegisterRoutes(router)

	protected := router.Group("/")

	protected.Use(middleware.NewHMACAuth(middleware.HMACConfig{
		Secret:         secret,
		MaxSkewSeconds: maxSkewSeconds,
	}))

	// authHandler.RegisterRoutes(protected)
	projectHandler.RegisterRoutes(protected)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
