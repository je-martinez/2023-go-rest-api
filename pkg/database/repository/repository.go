package repository

import "gorm.io/gorm"

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

func (r *GormRepository[T]) FindByID(id uint, preloads ...string) (*T, error) {
	var entity T
	err := r.DBWithPreloads(preloads).First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GormRepository[T]) FindByStringID(id string, preloads ...string) (*T, error) {
	var entity T
	err := r.DBWithPreloads(preloads).First(&entity, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GormRepository[T]) Find(query T, preloads ...string) (*T, error) {
	var entity T
	err := r.DBWithPreloads(preloads).Where(query).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GormRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *GormRepository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
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
