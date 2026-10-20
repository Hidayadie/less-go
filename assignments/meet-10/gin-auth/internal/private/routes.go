package private

import (
	"net/http"

	"gin-auth/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, requireAuth gin.HandlerFunc) {
	group := r.Group("/api/private")
	group.Use(requireAuth)
	group.GET("/profile", middleware.MockEndpoint("private.profile"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": c.MustGet("user")})
	})
}
