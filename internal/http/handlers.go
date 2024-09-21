package http

import (
	"github.com/askaroe/hexago/internal/domain/service"
	"github.com/askaroe/hexago/internal/handler"
	"github.com/askaroe/jsonlog"
)

type Handlers struct {
	user *handler.User
}

func newHandler(logger *jsonlog.Logger, services *service.Builder) Handlers {
	return Handlers{user: handler.NewUserHandler(&handler.Options{services, *logger})}
}
