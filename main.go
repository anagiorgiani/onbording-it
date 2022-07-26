package main

import (
	"database/sql"
	"log"
	"os"

	controller "github.com/anagiorgiani/onbording-it/cmd/server/controllers/my_api"
	db "github.com/anagiorgiani/onbording-it/db"
	"github.com/anagiorgiani/onbording-it/internal/my_api/domain"
	"github.com/anagiorgiani/onbording-it/internal/my_api/repository"
	"github.com/anagiorgiani/onbording-it/internal/my_api/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env", err)
	}

	storageDB := db.Init()
	server := gin.Default()

	apiRepository := buildRepositories(storageDB)

	funcHandlers(apiRepository, server)

	port := os.Getenv("HOST_PORT")
	server.Run(port)
}

func funcHandlers(apiRepository domain.ApiRepository, server *gin.Engine) {
	apiService := service.NewService(apiRepository)
	apiController := controller.NewController(apiService)

	ApiGroup := server.Group("/api/")
	ApiGroup.GET("/myApi", apiController.Get())
}

func buildRepositories(storageDB *sql.DB) domain.ApiRepository {

	apiRepository := repository.NewRepository(storageDB)
	return apiRepository
}
