package repo

import (
	"context"
	"github.com/askaroe/hexago/internal/adapter/database/entity"
	"github.com/askaroe/hexago/internal/domain/port/database"
	"github.com/askaroe/hexago/internal/domain/port/repository"
)

type user struct {
	db database.IDB
}

func NewUser(db database.IDB) repository.IUser {
	return &user{db: db}
}

func (u *user) GetAll(ctx context.Context) (res []entity.User, err error) {
	err = u.db.GetAll(ctx, res)
	return
}

func (u *user) GetById(ctx context.Context, id int64) (res entity.User, err error) {
	err = u.db.GetFirst(ctx, &res, id)
	return
}

func (u *user) Create(ctx context.Context, req entity.User) (err error) {
	err = u.db.Create(ctx, &req)
	return
}

func (u *user) Update(ctx context.Context, req entity.User) (err error) {
	err = u.db.Update(ctx, &req)
	return
}

func (u *user) Delete(ctx context.Context, req entity.User) (err error) {
	err = u.db.Delete(ctx, &req)
	return
}
