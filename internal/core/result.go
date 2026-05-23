package core

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result[T any] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func Success[T any](data T) *Result[T] {
	return &Result[T]{Code: "0", Message: "success", Data: data}
}

func Error(code string, message string) *Result[any] {
	return &Result[any]{Code: code, Message: message, Data: nil}
}

func ToResult(ctx *gin.Context, data any, err error) {
	if err == nil {
		ctx.JSON(http.StatusOK, Success(data))
		return
	}

	if ae, ok := errors.AsType[*AppError](err); ok {
		switch ae.Type {
		case InvalidRequestError:
			ctx.JSON(http.StatusBadRequest, Error(string(ae.Code), ae.Message))
		case InvalidArgument:
			ctx.JSON(http.StatusBadRequest, Error(string(ae.Code), ae.Message))
		case InternalServerError:
			ctx.JSON(http.StatusInternalServerError, Error(string(ae.Code), ae.Message))
		default:
			ctx.JSON(http.StatusInternalServerError, Error(string(InternalServerErrorCode), "system error"))
		}
		return
	}
	ctx.JSON(http.StatusInternalServerError, Error(string(InternalServerErrorCode), "system error"))
}
