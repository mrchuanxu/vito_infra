package dbhelper

import (
	"context"

	"github.com/pkg/errors"
)

// DBHelper 简化获取数据库db的操作
type DBHelper struct {
}

func (r *DBHelper) Create(ctx context.Context, m interface{}) error {
	db := DB(ctx)
	if db == nil {
		return errors.New("Sql object is nil")
	}

	return Create(db, m)
}

func (r *DBHelper) Save(ctx context.Context, m interface{}) error {
	db := DB(ctx)
	if db == nil {
		return errors.New("Sql object is nil")
	}

	return Save(db, m)
}

func (r *DBHelper) First(ctx context.Context, m interface{}, filter ...interface{}) error {
	db := DB(ctx)
	if db == nil {
		return errors.New("Sql object is nil")
	}

	return First(db, m, filter...)
}

func (r *DBHelper) Pluck(ctx context.Context, m interface{}, c string, v interface{}, filter ...interface{}) error {
	db := DB(ctx)
	if db == nil {
		return errors.New("Sql object is nil")
	}

	return Pluck(db, m, c, v, filter...)
}

func (r *DBHelper) Count(ctx context.Context, m, v interface{}, filter ...interface{}) error {
	db := DB(ctx)
	if db == nil {
		return errors.New("Sql object is nil")
	}
	db = db.Model(m)
	return Count(db, v, filter...)
}

func (r *DBHelper) Update(ctx context.Context, m interface{}) error {
	db := DB(ctx)
	if db == nil {
		return errors.New("Sql object is nil")
	}

	return Update(db, m)
}

func (r *DBHelper) Delete(ctx context.Context, m interface{}, filter ...interface{}) error {
	db := DB(ctx)
	if db == nil {
		return errors.New("Sql object is nil")
	}
	return Delete(db, m, filter...)
}

func (r *DBHelper) Find(ctx context.Context, m interface{}, filter ...interface{}) error {
	db := DB(ctx)
	if db == nil {
		return errors.New("Sql object is nil")
	}
	return Find(db, m, filter...)
}

func (r *DBHelper) Scan(ctx context.Context, v interface{}) error {
	db := DB(ctx)
	if db == nil {
		return errors.New("Sql object is nil")
	}
	return Scan(db, v)
}
