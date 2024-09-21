package repository

import (
	"context"
	"github.com/askaroe/hexago/internal/adapter/database/entity"
)

type Manager struct {
}

type IUser interface {
	GetAll(ctx context.Context) (res []entity.User, err error)
	GetById(ctx context.Context, id int64) (res entity.User, err error)
	Create(ctx context.Context, req entity.User) (err error)
	Update(ctx context.Context, req entity.User) (err error)
	Delete(ctx context.Context, req entity.User) (err error)
}
