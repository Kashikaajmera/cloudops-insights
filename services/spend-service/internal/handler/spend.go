
package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/kashika-ajmera/cloudops-insights/services/spend-service/internal/service"
)

type SpendHandler struct {
    service *service.SpendService
}

func NewSpendHandler(s *service.SpendService) *SpendHandler {
    return &SpendHandler{service: s}
}

func (h *SpendHandler) GetSpendSummary(c *gin.Context) {
    summary := h.service.GetSummary()
    c.JSON(http.StatusOK, summary)
}
