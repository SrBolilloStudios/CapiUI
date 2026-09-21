package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"capiui_pos/internal/models"
	"capiui_pos/internal/query"
)

type ProductStore struct {
	db *sql.DB
}

func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{db: db}
}

func (s *ProductStore) List(ctx context.Context, page, limit int) ([]models.Product, error) {
	if page < 1 || limit < 1 {
		return nil, errors.New("page and limit must be positive")
	}

	q, err := query.Get("products.list")
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	rows, err := s.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	return scanProducts(rows)
}

func (s *ProductStore) GetByIDs(ctx context.Context, ids []int) ([]models.Product, error) {
	if len(ids) == 0 {
		return nil, errors.New("at least one id is required")
	}

	q, err := query.Get("products.get_by_ids")

	placeholders, args := placeholders(ids)
	rows, err := s.db.QueryContext(ctx, fmt.Sprintf(q, strings.Join(placeholders, ",")), args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	return scanProducts(rows)
}

func (s *ProductStore) Create(ctx context.Context, products []models.Product) (int64, error) {
	if len(products) == 0 {
		return 0, errors.New("at least one product is required")
	}

	q, err := query.Get("products.insert")
	if err != nil {
		return 0, err
	}

	placeholders := make([]string, len(products))
	args := make([]any, 0, len(products)*7)
	for i, p := range products {
		placeholders[i] = "(?,?,?,?,?,?,?)"
		args = append(args, p.SKU, p.Name, p.Price, p.Description, p.Img, p.Stock, p.Category.ID)
	}

	result, err := s.db.ExecContext(ctx, fmt.Sprintf(q, strings.Join(placeholders, ",")), args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *ProductStore) Delete(ctx context.Context, ids []int) error {
	if len(ids) == 0 {
		return errors.New("at least one id is required")
	}

	q, err := query.Get("products.delete_by_ids")
	if err != nil {
		return err
	}

	placeholders, args := placeholders(ids)
	_, err = s.db.ExecContext(ctx, fmt.Sprintf(q, strings.Join(placeholders, ",")), args...)
	return err
}

func placeholders(ids []int) ([]string, []any) {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}
	return placeholders, args
}

func scanProducts(rows *sql.Rows) ([]models.Product, error) {
	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(
			&p.ID,
			&p.SKU,
			&p.Name,
			&p.Price,
			&p.Description,
			&p.Img,
			&p.Stock,
			&p.Category.ID,
			&p.Category.Name,
		); err != nil {
			return products, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}
