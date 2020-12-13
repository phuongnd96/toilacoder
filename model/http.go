package model

import "github.com/kataras/iris"

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(ctx iris.Context, status int, err error) {
	e := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	ctx.StatusCode(status)
	ctx.JSON(e)
}
