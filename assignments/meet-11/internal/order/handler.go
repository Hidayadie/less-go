package order

import (
	"fmt"
	"gin-auth/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var input struct {
		Total float64 `json:"total" gorm:"not null;default:1"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nth"})
		return
	}

	user := c.MustGet("user").(auth.User)
	order, err := h.service.CreateOrder(c.Request.Context(), user, input.Total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menyimpan order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": order})
}

func (h *Handler) CreateOrderDetail(c *gin.Context) {
	var input struct {
		OrderID   uint    `json:"order_id"`
		ProductID uint    `json:"product_id"`
		Qty       int     `json:"qty"`
		Price     float64 `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input tidak valid"})
		return
	}

	user := c.MustGet("user").(auth.User)
	orderDetail, err := h.service.CreateOrderDetail(c.Request.Context(), user, input.OrderID, input.ProductID, input.Qty, input.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menyimpan"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": orderDetail})
}

func (h *Handler) ListOrder(c *gin.Context) {
	user := c.MustGet("user").(auth.User)
	orders, err := h.service.GetOrderList(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "eror read"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (h *Handler) GetOrderByID(c *gin.Context) {
	user := c.MustGet("user").(auth.User)
	id := c.Param("id")
	var orderID uint
	_, err := fmt.Sscanf(id, "%d", &orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID order tidak valid"})
		return
	}
	order, err := h.service.GetOrderByID(c.Request.Context(), orderID, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}

