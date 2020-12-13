package router

import (
	"toilacoder/controller"

	"github.com/kataras/iris"
)

func FileRoutes(c *controller.Controller, api *iris.Application) {

	// api.Get("/ui", func(ctx iris.Context) {
	// 	auth, _ := constant.Sess.Start(ctx).GetBoolean("authenticated")
	// 	if !auth {
	// 		ctx.StatusCode(iris.StatusForbidden)
	// 		return
	// 	}
	// 	ctx.ViewLayout(iris.NoLayout)
	// 	ctx.View("file/upload_image.html")
	// })

	// api.Post("/upload-image", c.UploadImage)

	// api.Get("/list-image", func(ctx iris.Context) {
	// 	auth, _ := constant.Sess.Start(ctx).GetBoolean("authenticated")
	// 	if !auth {
	// 		ctx.StatusCode(iris.StatusForbidden)
	// 		return
	// 	}
	// 	var images []model.ListImage
	// 	_, err := c.DB.Query(&images, `
	// 		SELECT id, title, filename FROM image
	// 	`)
	// 	if err != nil {
	// 		log.Println(err)
	// 		ctx.JSON("Lỗi kho ảnh")
	// 		return
	// 	}
	// 	ctx.ViewData("images", images)
	// 	ctx.ViewLayout(iris.NoLayout)
	// 	ctx.View("file/list_image.html")
	// })

	// api.Get("/detail-image/{id}", func(ctx iris.Context) {
	// 	id := ctx.Params().Get("id")
	// 	auth, _ := constant.Sess.Start(ctx).GetBoolean("authenticated")
	// 	if !auth {
	// 		ctx.StatusCode(iris.StatusForbidden)
	// 		return
	// 	}
	// 	var image model.ListImage
	// 	_, err := c.DB.Query(&image, `
	// 		SELECT id, title, filename, source FROM image WHERE id = ?
	// 	`, id)
	// 	if err != nil {
	// 		log.Println(err)
	// 		ctx.JSON("Lỗi ảnh")
	// 		return
	// 	}
	// 	log.Println("-------------", image)
	// 	ctx.ViewData("image", image)
	// 	ctx.ViewLayout(iris.NoLayout)
	// 	ctx.View("file/detail_image.html")
	// })

	// api.Post("/delete-image", c.DeleteImage)
}
