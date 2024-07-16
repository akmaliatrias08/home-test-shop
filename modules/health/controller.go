package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func router(rg *gin.RouterGroup) {
	healthRoute := rg.Group("/health")

	healthRoute.GET("", healthCheck)
}

// CheckHealth      godoc
// @Summary     	Check health APIs
// @Description  	Responds if the APIs success.
// @Tags         	health
// @Produce      	json
// @Success      	200  {object}  health.HealthSuccess
// @Router       	/health [get]
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
