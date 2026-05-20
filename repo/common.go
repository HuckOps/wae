package repo

import (
	"context"
	"wae/db"

	"gorm.io/gorm"
)

type Option func(*gorm.DB) *gorm.DB

func WithWhere(where string, args ...any) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(where, args...)
	}
}

func GetByID[T any](ctx context.Context, id any) (T, error) {
	var model T
	err := db.Db.WithContext(ctx).Model(&model).Where("id = ?", id).First(&model).Error
	return model, err
}

func GetByOptions[T any](ctx context.Context, opts ...Option) ([]T, error) {
	var results []T
	var model T
	query := db.Db.WithContext(ctx).Model(&model)
	for _, opt := range opts {
		query = opt(query)
	}

	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func GetByPagination[T any](ctx context.Context, page, pageSize int, opts ...Option) (int64, []T, error) {
	var results []T
	var model T
	query := db.Db.WithContext(ctx).Model(&model)
	for _, opt := range opts {
		query = opt(query)
	}

	var count int64
	err := query.Count(&count).Error
	if err != nil {
		return 0, nil, err
	}

	query.Offset((page - 1) * pageSize)
	query.Limit(pageSize)

	err = query.Find(&results).Error
	return count, results, err
}

func Create[T any](ctx context.Context, model T) error {
	return db.Db.WithContext(ctx).Create(&model).Error
}
