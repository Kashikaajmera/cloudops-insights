
package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/kashika-ajmera/cloudops-insights/services/spend-service/internal/generated"
    "github.com/kashika-ajmera/cloudops-insights/services/spend-service/internal/handler"
    "github.com/kashika-ajmera/cloudops-insights/services/spend-service/internal/service"
)

func main() {
    router := gin.New()
    router.Use(gin.Logger(), gin.Recovery())

    spendService := service.NewSpendService()
    spendHandler := handler.NewSpendHandler(spendService)

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
