package main

import (
	"context"
	"database/sql"
	"log"

	"capiui_pos/internal/models"
	"capiui_pos/internal/query"
	"capiui_pos/internal/store"

	_ "modernc.org/sqlite"
)

func main() {
	ctx := context.Background()

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

	mustExec(db, "categories.drop_table")
	mustExec(db, "products.drop_table")
	mustExec(db, "categories.create_table")
	mustExec(db, "products.create_table")

	categories := []models.Category{
		{Name: "Keyword"},
		{Name: "Network"},
		{Name: "Display"},
	}
	categoriesCreated, err := store.NewCategoryStore(db).Create(ctx, categories)
	if err != nil {
		log.Fatalf("failed to insert categories: %v", err)
	}

	products := []models.Product{
		{
			Name:        "Teclado Mecánico RGB",
			Price:       89.99,
			Description: "Teclado con switches azules",
			Stock:       15,
			Img:         "/static/img/products.png",
			Category:    models.Category{ID: 1},
			SKU:         "TEC-001",
		},
		{
			Name:        "Ratón Inalámbrico",
			Price:       25.50,
			Description: "Ratón ergonómico recargable",
			Stock:       30,
			Img:         "/static/img/products.png",
			Category:    models.Category{ID: 2},
			SKU:         "RAT-002",
		},
		{
			Name:        "Monitor 24 pulgadas",
			Price:       150.00,
			Description: "Monitor IPS 144Hz",
			Stock:       10,
			Img:         "/static/img/products.png",
			Category:    models.Category{ID: 3},
			SKU:         "MON-003",
		},
	}

	created, err := store.NewProductStore(db).Create(ctx, products)
	if err != nil {
		log.Fatalf("failed to insert products: %v", err)
	}
	log.Printf("Seed completed: %d categories, %d products", categoriesCreated, created)
}

func mustExec(db *sql.DB, queryName string) {
	q, err := query.Get(queryName)
	if err != nil {
		log.Fatalf("failed to get query %s: %v", queryName, err)
	}
	if _, err := db.Exec(q); err != nil {
		log.Fatalf("failed to execute query %s: %v", queryName, err)
	}
}
