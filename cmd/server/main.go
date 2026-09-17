package main

import (
	"fmt"
	"log"
	"net/http"

	"capiui_pos/internal/api"
)

func main() {
	fmt.Print("Servidor inicializado")
	mux := http.NewServeMux()
	mux.HandleFunc("/api/test/products", api.ProductsHandler)
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(mux)))

}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
