package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MarcoVitangeli/admin-dash/models"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type ProductRepository interface {
	CreateCategory(ctx context.Context, name string) (int64, error)
	CreateProduct(ctx context.Context, name string, description string, categoryId int) (int64, error)
	GetCategories(ctx context.Context, count uint) ([]models.ProductCategory, error)
}

type productDB struct {
	*sql.DB
}

var (
	instance              ProductRepository
	ErrDuplicatedCategory = errors.New("category is already present")
	ErrDuplicatedProduct  = errors.New("product is already present")
	ErrCategoryNotFound   = errors.New("category not found: id nonexistent")
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

func (pdb *productDB) CreateProduct(ctx context.Context, name string, description string, categoryId int) (int64, error) {
	var (
		aux     int
		auxName string
	)
	if err := pdb.QueryRowContext(ctx, "SELECT 1 FROM product_category WHERE id = ?", categoryId).Scan(&aux); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, ErrCategoryNotFound
		}
		return -1, fmt.Errorf("database error: %w", err)
	}

	if err := pdb.QueryRowContext(ctx, "SELECT 1 FROM products WHERE name = ?", name).Scan(&auxName); err == nil {
		return -1, ErrDuplicatedProduct
	} else if !errors.Is(err, sql.ErrNoRows) {
		return -1, fmt.Errorf("unexpected error in db: %w", err)
	}

	res, err := pdb.ExecContext(ctx, "INSERT INTO products(name, description, category_id) VALUES (?,?,?)",
		name,
		description,
		categoryId)

	if err != nil {
		return -1, fmt.Errorf("unexpected failure creating product: %w", err)
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("unexpected failure getting the created product ID: %w", err)
	}

	return newId, nil
}

func (pdb *productDB) GetCategories(ctx context.Context, count uint) ([]models.ProductCategory, error) {
	rows, err := pdb.QueryContext(ctx, "SELECT id, name, created_at FROM product_category LIMIT ?", count)
	if err != nil {
		return nil, fmt.Errorf("error reading categories: %w", err)
	}

    categories := make([]models.ProductCategory, 0, count)

	for rows.Next() {
		var (
			id       uint64 
			name      string
			createdAt time.Time
		)
        if err := rows.Scan(&id, &name, &createdAt); err != nil {
            return nil, fmt.Errorf("error reading row: %w", err)
        }
        categories = append(categories, models.ProductCategory{
            Id: id,
            Name: name,
            CreatedAt: createdAt,
        })
	}

    return categories[:len(categories):len(categories)], nil
}
