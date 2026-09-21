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

type CategoryStore struct {
	db *sql.DB
}

func NewCategoryStore(db *sql.DB) *CategoryStore {
	return &CategoryStore{db: db}
}

func (s *CategoryStore) List(ctx context.Context, page, limit int) ([]models.Category, error) {
	if page < 1 || limit < 1 {
		return nil, errors.New("page and limit must be positive")
	}

	q, err := query.Get("categories.list")
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

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return categories, err
		}
		categories = append(categories, c)
	}
	return categories, rows.Err()
}

func (s *CategoryStore) GetByIDs(ctx context.Context, ids []int) ([]models.Category, error) {
	if len(ids) == 0 {
		return nil, errors.New("at least one id is required")
	}

	q, err := query.Get("categories.get_by_ids")
	if err != nil {
		return nil, err
	}

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

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return categories, err
		}
		categories = append(categories, c)
	}
	return categories, rows.Err()
}

func (s *CategoryStore) Create(ctx context.Context, categories []models.Category) (int64, error) {
	if len(categories) == 0 {
		return 0, errors.New("at least one category is required")
	}

	q, err := query.Get("categories.insert")
	if err != nil {
		return 0, err
	}

	placeholders := make([]string, len(categories))
	args := make([]any, len(categories))
	for i, c := range categories {
		placeholders[i] = "(?)"
		args[i] = c.Name
	}

	result, err := s.db.ExecContext(ctx, fmt.Sprintf(q, strings.Join(placeholders, ",")), args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *CategoryStore) Delete(ctx context.Context, ids []int) error {
	if len(ids) == 0 {
		return errors.New("at least one id is required")
	}

	q, err := query.Get("categories.delete_by_ids")
	if err != nil {
		return err
	}

	placeholders, args := placeholders(ids)
	_, err = s.db.ExecContext(ctx, fmt.Sprintf(q, strings.Join(placeholders, ",")), args...)
	return err
}
