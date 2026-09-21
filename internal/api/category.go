package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"capiui_pos/internal/models"
	"capiui_pos/internal/store"
)

type CategoryHandler struct {
	categories *store.CategoryStore
}

func NewCategoryHandler(categories *store.CategoryStore) *CategoryHandler {
	return &CategoryHandler{categories: categories}
}

func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	page, limit, ok := parsePagination(w, r)
	if !ok {
		return
	}

	categories, err := h.categories.List(r.Context(), page, limit)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list categories")
		return
	}
	writeJSON(w, http.StatusOK, categories)
}

func (h *CategoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	ids, err := parseIDs(r.PathValue("ids"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid category id")
		return
	}

	categories, err := h.categories.GetByIDs(r.Context(), ids)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to get categories")
		return
	}
	if len(categories) == 0 {
		writeError(w, http.StatusNotFound, "categories not found")
		return
	}
	writeJSON(w, http.StatusOK, categories)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	if err := json.NewDecoder(r.Body).Decode(&categories); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if err := validateCategories(categories); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	created, err := h.categories.Create(r.Context(), categories)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create categories")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int64{"created": created})
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ids, err := parseIDs(r.PathValue("ids"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid category id")
		return
	}

	if err := h.categories.Delete(r.Context(), ids); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to delete categories")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func validateCategories(categories []models.Category) error {
	if len(categories) == 0 {
		return errors.New("at least one category is required")
	}
	for _, c := range categories {
		if c.Name == "" {
			return errors.New("category name is required")
		}
	}
	return nil
}
