package repositories

import (
	"time"
	"gorm.io/gorm"
	"context"
)

type IBaseRepository [T any] interface {
	Create (ctx context.Context, t T) (T, error)
	GetById (ctx context.Context, id int64) (T, error)
	GetAll (ctx context.Context) ([]T, error)
	Update (ctx context.Context, t T) (T, error)
}

type BaseRepository [T any] struct {
	Db *gorm.DB
}

func (baseRepository BaseRepository[T]) GetById (ctx context.Context, id int64) (T, error) {
	var t T 
	time.Sleep(3 * time.Second)
	err := baseRepository.Db.WithContext(ctx).First(&t, id).Error
  	return t, err
}

func (baseRepository BaseRepository[T]) GetAll (ctx context.Context) ([]T, error) {
	var t []T
	time.Sleep(3 * time.Second)
	err := baseRepository.Db.WithContext(ctx).Find(&t).Error
  	return t, err
}

func (baseRepository BaseRepository [T]) Create (ctx context.Context, t T) (T, error) {
	err := baseRepository.Db.WithContext(ctx).Create(&t).Error
	return t, err
}

func (baseRepository BaseRepository [T]) Update (ctx context.Context, t T) (T, error) {
	err := baseRepository.Db.WithContext(ctx).Save(&t).Error
	return t, err
}