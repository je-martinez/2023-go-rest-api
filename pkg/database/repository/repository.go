package repositories

import (
	"context"
	e "main/pkg/database/extensions"

	"gorm.io/gorm"
)

type GormModel[M any] interface {
	ToModel() M
	FromModel(entity M) interface{}
}

func NewRepository[E GormModel[M], M any](db *gorm.DB) *GormRepository[E, M] {
	return &GormRepository[E, M]{
		db: db,
	}
}

type GormRepository[E GormModel[M], M any] struct {
	db *gorm.DB
}

func (r *GormRepository[E, M]) Insert(ctx context.Context, db_model *M) (*E, error) {
	var start E
	entity := start.FromModel(*db_model).(E)

	err := r.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *GormRepository[E, M]) Delete(ctx context.Context, db_model *M) error {
	var start E
	entity := start.FromModel(*db_model).(E)
	err := r.db.WithContext(ctx).Delete(entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *GormRepository[E, M]) DeleteById(ctx context.Context, id any) error {
	var start E
	err := r.db.WithContext(ctx).Delete(&start, &id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *GormRepository[E, M]) Update(ctx context.Context, db_model *M) (*E, error) {
	var start E
	entity := start.FromModel(*db_model).(E)

	err := r.db.WithContext(ctx).Save(&entity).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *GormRepository[E, M]) FindByID(ctx context.Context, id any) (E, error) {
	var entity E
	err := r.db.WithContext(ctx).First(&entity, id).Error
	if err != nil {
		return *new(E), err
	}

	return entity, nil
}

func (r *GormRepository[E, M]) Find(ctx context.Context, specifications ...e.Specification) ([]E, error) {
	return r.FindWithLimit(ctx, -1, -1, specifications...)
}

func (r *GormRepository[E, M]) Count(ctx context.Context, specifications ...e.Specification) (i int64, err error) {
	model := new(E)
	err = r.getPreWarmDbForSelect(ctx, specifications...).Model(model).Count(&i).Error
	return
}

func (r *GormRepository[E, M]) getPreWarmDbForSelect(ctx context.Context, specification ...e.Specification) *gorm.DB {
	dbPrewarm := r.db.WithContext(ctx)
	for _, s := range specification {
		dbPrewarm = dbPrewarm.Where(s.GetQuery(), s.GetValues()...)
	}
	return dbPrewarm
}

func (r *GormRepository[E, M]) FindWithLimit(ctx context.Context, limit int, offset int, specifications ...e.Specification) ([]E, error) {
	var entities []E

	dbPrewarm := r.getPreWarmDbForSelect(ctx, specifications...)
	err := dbPrewarm.Limit(limit).Offset(offset).Find(&entities).Error

	if err != nil {
		return nil, err
	}

	result := make([]E, 0, len(entities))
	for _, row := range entities {
		result = append(result, row)
	}

	return result, nil
}

func (r *GormRepository[E, M]) FindAll(ctx context.Context) ([]E, error) {
	return r.FindWithLimit(ctx, -1, -1)
}
