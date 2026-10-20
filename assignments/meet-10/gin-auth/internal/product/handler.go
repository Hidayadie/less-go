package product

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"gin-auth/internal/auth"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama product wajib diisi, maksimal 100 karakter"})
		return
	}
	user := c.MustGet("user").(auth.User)
	item, err := h.service.Create(c.Request.Context(), input.Name, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan product"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": item})
}

func (h *Handler) List(c *gin.Context) {
	user := c.MustGet("user").(auth.User)
	items, err := h.service.List(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h *Handler) ListByID(c *gin.Context) {
	productID, ok := productIDParam(c)
	if !ok {
		return
	}
	user := c.MustGet("user").(auth.User)
	item, err := h.service.ListByID(c.Request.Context(), productID, user)
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func (h *Handler) Put(c *gin.Context) {
	productID, ok := productIDParam(c)
	if !ok {
		return
	}
	var input struct {
		Name string `json:"name" binding:"required,max=100"`
	}
	if err := c.ShouldBindJSON(&input); err != nil || strings.TrimSpace(input.Name) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama product wajib diisi, maksimal 100 karakter"})
		return
	}
	user := c.MustGet("user").(auth.User)
	items, err := h.service.Put(c.Request.Context(), productID, Product{Name: strings.TrimSpace(input.Name), UserID: user.ID})
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h *Handler) Delete(c *gin.Context) {
	productID, ok := productIDParam(c)
	if !ok {
		return
	}
	user := c.MustGet("user").(auth.User)
	if err := h.service.Delete(c.Request.Context(), productID, user); err != nil {
		if errors.Is(err, ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"id": productID, "deleted": true}})
}

// productIDParam membaca parameter :id dan memastikan berupa angka positif.
func productIDParam(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID product tidak valid"})
		return 0, false
	}
	return uint(id), true
}
