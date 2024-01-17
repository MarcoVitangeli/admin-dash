package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type ProductRepository interface {
	CreateCategory(ctx context.Context, name string) (int64, error)
}

type productDB struct {
	*sql.DB
}

var (
	instance              ProductRepository
	ErrDuplicatedCategory = errors.New("Error: category is alreacy present")
)

func GetRepository() (ProductRepository, error) {
	if instance != nil {
		return instance, nil
	}

	db, err := sql.Open("sqlite3", "./products.db")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	instance = &productDB{db}
	return instance, nil
}

func (pdb *productDB) CreateCategory(ctx context.Context, name string) (int64, error) {
	var (
		err error
		aux int
	)
	if err = pdb.QueryRowContext(ctx, "SELECT 1 FROM product_category WHERE name = ?", name).Scan(&aux); err == nil {
		return -1, ErrDuplicatedCategory
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return -1, fmt.Errorf("Error reading row: %w", err)
	}

	if res, err := pdb.ExecContext(ctx, "INSERT INTO product_category(name) VALUES (?)", name); err != nil {
		return -1, fmt.Errorf("error inserting category in database: %w", err)
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return -1, fmt.Errorf("error checking last category id: %w", err)
		}
		return id, nil
	}
}
