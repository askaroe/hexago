package service

import (
	"github.com/askaroe/hexago/internal/adapter/database/repo"
	"github.com/askaroe/hexago/internal/domain/port/iservice"
	"github.com/askaroe/hexago/internal/domain/service/user"
)

type Builder struct {
	User iservice.IUser
}

func Init(db *repo.Builder) *Builder {
	return &Builder{User: user.New(db)}
}
