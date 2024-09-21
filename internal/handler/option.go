package handler

import (
	"github.com/askaroe/hexago/internal/domain/service"
	"github.com/askaroe/jsonlog"
)

type Options struct {
	Services *service.Builder
	Logger   jsonlog.Logger
}
