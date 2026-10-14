package analytic

import (
	"encoding/json"
	"net/http"
	"log"
	"context"
)

// handler.go
type Handler struct{ Service Service }

func (h Handler) Customers(w http.ResponseWriter, r *http.Request) {
	h.serve(w, r, h.Service.Customers)
}
func (h Handler) Products(w http.ResponseWriter, r *http.Request) {
	h.serve(w, r, h.Service.Products)
}

func (h Handler) serve(w http.ResponseWriter, r *http.Request,
	fetch func(context.Context) (json.RawMessage, error)) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "method tidak didukung", http.StatusMethodNotAllowed)
		return
	}
	rows, err := fetch(r.Context())
	if err != nil {
		log.Printf("analytics: %v", err) // di summary, error asli tidak pernah dicatat
		http.Error(w, "query database gagal", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write(rows)
}
