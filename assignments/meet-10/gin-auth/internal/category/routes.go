package category

import (
	"gin-auth/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup, h * Handler) {
	group.POST("/category", middleware.MockEndpoint("private.category.create"), h.Create)
	group.GET("/category", middleware.MockEndpoint("private.category.list"), h.List)
}
