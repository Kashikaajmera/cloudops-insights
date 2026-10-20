package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kashika-ajmera/cloudops-insights/services/spend-service/internal/config"
	"github.com/kashika-ajmera/cloudops-insights/services/spend-service/internal/generated"
	"github.com/kashika-ajmera/cloudops-insights/services/spend-service/internal/handler"
	"github.com/kashika-ajmera/cloudops-insights/services/spend-service/internal/repository"
	"github.com/kashika-ajmera/cloudops-insights/services/spend-service/internal/service"
)

func main() {
	ctx := context.Background()

	dbPool, err := config.NewDatabasePool(ctx)
	if err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}
	defer dbPool.Close()

	spendRepository := repository.NewPostgresRepository(dbPool)
	spendService := service.NewSpendService(spendRepository)
	spendHandler := handler.NewSpendHandler(spendService)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "spend-service",
		})
	})

	generated.RegisterHandlers(router, spendHandler)

	log.Println("Spend Service starting on port 8081")

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
