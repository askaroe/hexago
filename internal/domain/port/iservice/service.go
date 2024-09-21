package iservice

import "context"

type IUser interface {
	GetById(ctx context.Context, Id int64) (res interface{}, err error)
	GetAll(ctx context.Context) (res interface{}, err error)
}
