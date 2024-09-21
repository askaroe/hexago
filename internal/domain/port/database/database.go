package database

import "context"

type IDB interface {
	Close() (err error)
	GetAll(ctx context.Context, dest interface{}, conds ...interface{}) (err error)
	GetFirst(ctx context.Context, dest interface{}, conds ...interface{}) (err error)
	Create(ctx context.Context, value interface{}) (err error)
	Update(ctx context.Context, value interface{}) (err error)
	Delete(ctx context.Context, value interface{}) (err error)
}
