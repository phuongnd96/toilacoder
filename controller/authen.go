package controller

import (
	"toilacoder/constant"
	"toilacoder/model"

	"github.com/kataras/iris"
)

func (c *Controller) Login(ctx iris.Context) {
	var req model.Login
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Không hợp lệ.")
		return
	}

	if req.Username == "" || req.Passwd == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Username và Password không được để trống")
		return
	}

	if req.Passwd == constant.Account[req.Username] {
		session := constant.Sess.Start(ctx)
		session.Set("authenticated", true)
	} else {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Không hợp lệ.")
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Yêu cầu thành công")
}

func (c *Controller) Logout(ctx iris.Context) {
	session := constant.Sess.Start(ctx)

	// Revoke users authentication
	session.Set("authenticated", false)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Yêu cầu thành công")
}
