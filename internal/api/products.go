package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		fmt.Println("Metodo POST detectado")
		products := []Product{
			{ID: 1, Name: "Zanahoria", Price: 50.99},
			{ID: 2, Name: "Mango", Price: 33.2},
			{ID: 3, Name: "Limon", Price: 10.1},
		}
		json.NewEncoder(w).Encode(products)
	case http.MethodPost:
		var newProduct Product
		fmt.Println("Metodo GET detectado")
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "JSON invalido", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newProduct)
	default:
		http.Error(w, "Metodo no permitido,", http.StatusMethodNotAllowed)
	}
}
