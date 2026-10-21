package order

import (
	"gin-auth/internal/middleware"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(group *gin.RouterGroup, h *Handler) {
	group.POST("/order", middleware.MockEndpoint("private.order.create"), h.CreateOrder)
	group.POST("/orderDetail", middleware.MockEndpoint("private.orderdetail.create"), h.CreateOrderDetail)
	group.GET("/orders", middleware.MockEndpoint("private.order.list"), h.ListOrder)
	group.GET("/order/:id", middleware.MockEndpoint("private.order.get"), h.GetOrderByID)
 //	group.GET("/orderDetail", middleware.MockEndpoint("private.orderdetail.list"), h.Get

}
