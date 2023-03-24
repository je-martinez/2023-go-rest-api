package repository

import "gorm.io/gorm"

type GormRepository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db}
}

func (r *GormRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *GormRepository[T]) FindByID(id uint) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GormRepository[T]) FindByStringID(id string) (*T, error) {
	var entity T
	err := r.db.First(&entity, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GormRepository[T]) Find(query T) (*T, error) {
	var entity T
	err := r.db.Where(query).First(&entity).Error
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
