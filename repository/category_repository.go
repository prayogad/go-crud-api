package repository

import (
	"belajar_golang_restfult_api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, Tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, Tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, Tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, Tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, Tx *sql.Tx) []domain.Category
}
