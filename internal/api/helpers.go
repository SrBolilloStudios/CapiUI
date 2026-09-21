package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func parseIDs(raw string) ([]int, error) {
	parts := strings.Split(raw, ",")
	ids := make([]int, 0, len(parts))
	for _, part := range parts {
		id, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, fmt.Errorf("invalid id %q", part)
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func parsePositiveInt(raw string, fallback int) (int, error) {
	if raw == "" {
		return fallback, nil
	}
	n, err := strconv.Atoi(raw)
	if err != nil || n < 1 {
		return 0, fmt.Errorf("invalid value %q", raw)
	}
	return n, nil
}

// parsePagination reads the page and limit query parameters, writing a 400
// response and returning ok=false when either is invalid.
func parsePagination(w http.ResponseWriter, r *http.Request) (page, limit int, ok bool) {
	page, err := parsePositiveInt(r.URL.Query().Get("page"), defaultPage)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid page parameter")
		return 0, 0, false
	}
	limit, err = parsePositiveInt(r.URL.Query().Get("limit"), defaultLimit)
	if err != nil || limit > maxLimit {
		writeError(w, http.StatusBadRequest, "invalid limit parameter")
		return 0, 0, false
	}
	return page, limit, true
}
