package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"capiui_pos/internal/models"
	"capiui_pos/internal/store"
)

const (
	defaultPage  = 1
	defaultLimit = 50
	maxLimit     = 100
)

type ProductHandler struct {
	products *store.ProductStore
}

func NewProductHandler(products *store.ProductStore) *ProductHandler {
	return &ProductHandler{products: products}
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	page, limit, ok := parsePagination(w, r)
	if !ok {
		return
	}

	products, err := h.products.List(r.Context(), page, limit)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list products")
		return
	}
	writeJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	ids, err := parseIDs(r.PathValue("ids"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid product id")
		return
	}

	products, err := h.products.GetByIDs(r.Context(), ids)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to get products")
		return
	}
	if len(products) == 0 {
		writeError(w, http.StatusNotFound, "products not found")
		return
	}
	writeJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	if err := json.NewDecoder(r.Body).Decode(&products); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if err := validateProducts(products); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	created, err := h.products.Create(r.Context(), products)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create products")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int64{"created": created})
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ids, err := parseIDs(r.PathValue("ids"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid product id")
		return
	}

	if err := h.products.Delete(r.Context(), ids); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to delete products")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func validateProducts(products []models.Product) error {
	if len(products) == 0 {
		return errors.New("at least one product is required")
	}
	for _, p := range products {
		if p.Name == "" {
			return errors.New("product name is required")
		}
		if p.SKU == "" {
			return errors.New("product sku is required")
		}
		if p.Category.ID < 1 {
			return errors.New("product category is required")
		}
	}
	return nil
}
