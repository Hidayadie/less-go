package category

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required,max=100"`
	}
	if err := c.ShouldBindJSON(&input); err != nil || strings.TrimSpace(input.Name) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama kategori wajib diisi, maksimal 100 karakter"})
		return
	}
	category, err := h.service.Create(c.Request.Context(), input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan kategori"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": category})
}

func (h *Handler) List(c *gin.Context) {
	categories, err := h.service.Get(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca kategori"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}
