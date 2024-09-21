package repo

import (
	"github.com/askaroe/hexago/internal/domain/port/database"
	"github.com/askaroe/hexago/internal/domain/port/repository"
)

type Builder struct {
	User repository.IUser
}

func Init(db database.IDB) *Builder {
	return &Builder{User: NewUser(db)}
}
