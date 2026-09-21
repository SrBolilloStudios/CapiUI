package api

import "net/http"

// NewHandler registers every API route and returns the root HTTP handler.
func NewHandler(products *ProductHandler, categories *CategoryHandler) http.Handler {
	mux := http.NewServeMux()
	registerProductRoutes(mux, products)
	registerCategoryRoutes(mux, categories)
	return CORS(mux)
}

func registerProductRoutes(mux *http.ServeMux, h *ProductHandler) {
	mux.HandleFunc("GET /api/products", h.List)
	mux.HandleFunc("POST /api/products", h.Create)
	mux.HandleFunc("GET /api/products/{ids}", h.Get)
	mux.HandleFunc("DELETE /api/products/{ids}", h.Delete)
}

func registerCategoryRoutes(mux *http.ServeMux, h *CategoryHandler) {
	mux.HandleFunc("GET /api/categories", h.List)
	mux.HandleFunc("POST /api/categories", h.Create)
	mux.HandleFunc("GET /api/categories/{ids}", h.Get)
	mux.HandleFunc("DELETE /api/categories/{ids}", h.Delete)
}
