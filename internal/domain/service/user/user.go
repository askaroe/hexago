package user

import (
	"context"
	"github.com/askaroe/hexago/internal/adapter/database/repo"
	"github.com/askaroe/hexago/internal/domain/port/iservice"
	"github.com/askaroe/hexago/internal/domain/port/repository"
)

type userSrv struct {
	repo repository.IUser
}

func New(repo *repo.Builder) iservice.IUser {
	return &userSrv{repo: repo.User}
}

func (u *userSrv) GetById(ctx context.Context, Id int64) (res interface{}, err error) {
	res, err = u.repo.GetById(ctx, Id)

	return
}

func (u *userSrv) GetAll(ctx context.Context) (res interface{}, err error) {
	res, err = u.repo.GetAll(ctx)

	return
}
