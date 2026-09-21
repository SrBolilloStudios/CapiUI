package main

import (
	"database/sql"
	"log"
	"net/http"

	"capiui_pos/internal/api"
	"capiui_pos/internal/query"
	"capiui_pos/internal/store"

	_ "modernc.org/sqlite"
)

func main() {
	if err := query.Load(); err != nil {
		log.Fatalf("failed to load queries: %v", err)
	}

	db, err := sql.Open("sqlite", "./products.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close database: %v", err)
		}
	}()

	handler := api.NewHandler(
		api.NewProductHandler(store.NewProductStore(db)),
		api.NewCategoryHandler(store.NewCategoryStore(db)),
	)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
