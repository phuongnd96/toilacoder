package router

import (
	"toilacoder/controller"

	"github.com/kataras/iris"
)

func BasicRoutes(c *controller.Controller, api *iris.Application) {
	api.Post("/login", c.Login)
	api.Post("/confirm-the-notification", c.ConfirmTheNotification)
	api.Get("/login", func(ctx iris.Context) {
		ctx.ViewLayout(iris.NoLayout)
		ctx.View("login.html")
	})
	api.Get("/logout", c.Logout)

	// api.Get("/lien-he", func(ctx iris.Context) {
	// 	ctx.ViewData("ogURL", constant.DOMAIN+"/lien-he")
	// 	ctx.ViewData("title", "Liên hệ - Hướng dẫn, video, highlight, montage")
	// 	ctx.ViewData("ogTitle", "Liên hệ - Hướng dẫn, video, highlight, montage")
	// 	ctx.ViewData("ogImage", constant.DOMAIN+"/resources/image/lien-minh-huyen-thoai-toc-chien-logo.png")
	// 	ctx.ViewData("ogDescription", "Liên hệ - Trang web hướng dẫn, tin tức, highlight, montage, chia sẻ, funny, lên đồ các tựa game Liên Minh Huyền Thoại, Lol Mobile Tốc Chiến, Liên Quân Mobile, Minecraft, Đế chế")
	// 	ctx.ViewData("metaDescription", "Liên hệ - Trang web hướng dẫn, tin tức, highlight, montage, chia sẻ, funny, lên đồ các tựa game Liên Minh Huyền Thoại, Lol Mobile Tốc Chiến, Liên Quân Mobile, Minecraft, Đế chế")
	// 	ctx.View("contact.html")
	// })

	// api.Get("/thank-you", func(ctx iris.Context) {
	// 	ctx.ViewData("ogURL", constant.DOMAIN+"/thank-you")
	// 	ctx.ViewData("title", "Cảm ơn - Hướng dẫn, video, highlight, montage")
	// 	ctx.ViewData("ogTitle", "Cảm ơn - Hướng dẫn, video, highlight, montage")
	// 	ctx.ViewData("ogImage", constant.DOMAIN+"/resources/image/lien-minh-huyen-thoai-toc-chien-logo.png")
	// 	ctx.ViewData("ogDescription", "Trang cảm ơn - Lol Moblie Tốc Chiến - Hướng dẫn, video, highlight, montage")
	// 	ctx.ViewData("metaDescription", "Trang cảm ơn - Lol Moblie Tốc Chiến - Hướng dẫn, video, highlight, montage")
	// 	ctx.View("tks.html")
	// })

	api.Get("/404", func(ctx iris.Context) {
		ctx.ViewLayout(iris.NoLayout)
		ctx.StatusCode(iris.StatusOK)
		ctx.View("error/error.html")
	})
}
