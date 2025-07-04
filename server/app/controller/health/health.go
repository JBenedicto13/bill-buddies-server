package health

import (
	"billbuddies/server/app/helper"
	"billbuddies/server/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController handles health-related endpoints
type HealthController struct{}

// NewHealthController initializes a new HealthController
func NewHealthController() *HealthController {
	return &HealthController{}
}

// Check processes a health check request
// @Summary Health check
// @Description Get the health status of the server
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func (hc *HealthController) Check(c *gin.Context) {
	appConfig := config.LoadAppConfig()

	// Respond with success
	helper.FormatResponse(c, "success", http.StatusOK, "Health check successful", gin.H{
		"name":    appConfig.AppName,
		"version": appConfig.AppVersion,
	}, nil)
}