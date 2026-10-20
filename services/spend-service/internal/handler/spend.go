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
	// Temporary, server-controlled development tenant.
	// Replace with verified tenant context after authentication.
	const organizationID int64 = 1

	summary, err := h.service.GetSummary(
		c.Request.Context(),
		organizationID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve spend summary",
		})
		return
	}

	c.JSON(http.StatusOK, summary)
}
