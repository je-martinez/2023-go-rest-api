package repository

import (
	"errors"

	"github.com/je-martinez/2023-go-rest-api/pkg/types"

	"gorm.io/gorm"
)

type GormRepository[T any] struct {
	db            *gorm.DB
	default_joins []string
}

func NewRepository[T any](db *gorm.DB, defaultJoins []string) *GormRepository[T] {
	return &GormRepository[T]{db, defaultJoins}
}

func (r *GormRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *GormRepository[T]) FindByID(id uint, preloads ...string) (*T, bool, error) {
	var entity T
	err := r.DBWithPreloads(preloads).First(&entity, id).Error
	if err != nil {
		return nil, errors.Is(err, gorm.ErrRecordNotFound), err
	}
	return &entity, false, nil
}

func (r *GormRepository[T]) FindByStringID(id string, preloads ...string) (*T, bool, error) {
	var entity T
	err := r.DBWithPreloads(preloads).First(&entity, "id = ?", id).Error
	if err != nil {
		return nil, errors.Is(err, gorm.ErrRecordNotFound), err
	}
	return &entity, false, nil
}

func (r *GormRepository[T]) Find(options types.QueryOptions) (*T, bool, error) {
	var entity T
	err := r.DBWithPreloads(options.Preloads).Where(options.Query, options.Args).First(&entity).Error
	if err != nil {
		return nil, errors.Is(err, gorm.ErrRecordNotFound), err
	}
	return &entity, false, nil
}

func (r *GormRepository[T]) Update(entity *T) (bool, error) {
	err := r.db.Save(entity).Error
	return errors.Is(err, gorm.ErrRecordNotFound), err
}

func (r *GormRepository[T]) Delete(entity *T) (bool, error) {
	err := r.db.Delete(entity).Error
	return errors.Is(err, gorm.ErrRecordNotFound), err
}

func (r *GormRepository[T]) DBWithPreloads(preloads []string) *gorm.DB {
	dbConn := r.db

	for _, join := range r.default_joins {
		dbConn = dbConn.Joins(join)
	}

	for _, preload := range preloads {
		dbConn = dbConn.Preload(preload)
	}

	return dbConn
}
