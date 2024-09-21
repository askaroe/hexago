package handler

import (
	"github.com/askaroe/hexago/common"
	"github.com/askaroe/hexago/internal/domain/port/iservice"
	"github.com/askaroe/jsonlog"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	logger  jsonlog.Logger
	userSrv iservice.IUser
}

func NewUserHandler(opt *Options) *User {
	return &User{logger: opt.Logger, userSrv: opt.Services.User}
}

func (h *User) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		h.logger.PrintError(err, nil)
		return
	}

	res, err := h.userSrv.GetById(ctx, id)

	if err != nil {
		h.logger.PrintError(err, nil)
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse{
		Code:    0,
		Message: "success",
		Data:    res,
	})
}

func (h *User) GetAll(ctx *gin.Context) {
	res, err := h.userSrv.GetAll(ctx)

	if err != nil {
		h.logger.PrintError(err, nil)
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse{
		Code:    0,
		Message: "success",
		Data:    res,
	})
}
