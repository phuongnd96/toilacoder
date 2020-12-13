package controller

import (
	"log"
	"net/http"
	"toilacoder/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg"
	"github.com/kataras/iris"
)

type Controller struct {
	// DB instance
	DB *pg.DB

	// Configuration
	Config config.Config
}

// UserClaim chứa thông tin lưu trong JWT
type UserClaim struct {
	UserId         string
	EmailConfirmed bool
	UserRoles      []int32
	jwt.StandardClaims
}

func NewController() *Controller {
	var c Controller
	return &c
}

func HanlerErr(ctx iris.Context, statusCode int32, message string) {
	ctx.ViewData("message", message)
	ctx.ViewData("status_code", statusCode)
	ctx.View("error/error.html")
}

func (c *Controller) WrapRouter(ctx iris.Context) {
	log.Println(ctx.GetCurrentRoute())
	ctx.SetCookie(
		&http.Cookie{
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		},
	)
}
