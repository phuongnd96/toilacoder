package router

import (
	"fmt"
	"log"
	"toilacoder/constant"
	"toilacoder/controller"
	"toilacoder/model"

	"github.com/kataras/iris"
)

func PostRoutes(c *controller.Controller, api *iris.Application) {
	api.Get("/", c.GetPosts)
	api.Get("/posts", c.GetPosts)
	api.Get("/{slug}", c.GetPostBySlug)
	api.Get("/series/{slug}", c.GetSeriesBySlug)

	api.Get("/cp", func(ctx iris.Context) {
		// auth, _ := constant.Sess.Start(ctx).GetBoolean("authenticated")
		// if !auth {
		// 	ctx.Redirect("/login")
		// 	return
		// }

		var series []model.GetSeries
		query := fmt.Sprintf(`
			SELECT series.id ,series.title
			FROM toilacoder.series 
		`)
		_, err := c.DB.Query(&series, query)
		if err != nil {
			log.Println(err)
			ctx.JSON("Lỗi phía server")
			return
		}
		ctx.ViewData("series", series)
		ctx.ViewLayout(iris.NoLayout)
		ctx.View("blog/create.html")
	})

	api.Get("/cs", func(ctx iris.Context) {
		auth, _ := constant.Sess.Start(ctx).GetBoolean("authenticated")
		if !auth {
			ctx.Redirect("/login")
			return
		}

		ctx.ViewLayout(iris.NoLayout)
		ctx.View("series/create.html")
	})

	api.Get("/up/{slug}", func(ctx iris.Context) {
		ctx.ViewLayout(iris.NoLayout)
		// Check if user is authenticated
		// auth, _ := constant.Sess.Start(ctx).GetBoolean("authenticated")
		// if !auth {
		// 	ctx.Redirect("/login")
		// 	return
		// }

		slug := ctx.Params().Get("slug")
		// category := ctx.Params().Get("category")

		var post model.PostDetail
		_, err := c.DB.Query(&post, `
			SELECT post.id, post.slug, post.description, post.content, post.view_count, post.author_id,
			post.thumbnail, post.title, post.tag as tag,
			to_char(post.created_at, 'DD/MM/YYYY') created_at
			FROM toilacoder.post post
			WHERE slug = ?
		`, slug)
		if err != nil {
			log.Println(err)
			ctx.JSON("Bài viết không tồn tại")
			return
		}

		if post.Id == 0 {
			ctx.JSON("Bài viết không tồn tại")
			return
		}

		var tagStr string
		for i, x := range post.Tag {
			if i == 0 {
				tagStr = x
			} else {
				tagStr = tagStr + "," + x
			}
		}

		var series []model.GetSeries
		query := fmt.Sprintf(`
			SELECT series.id ,series.title
			FROM toilacoder.series 
		`)
		_, err = c.DB.Query(&series, query)
		if err != nil {
			log.Println(err)
			ctx.JSON("Lỗi phía server")
			return
		}
		ctx.ViewData("series", series)

		ctx.ViewData("post", post)
		ctx.ViewData("tag", tagStr)
		ctx.View("blog/update.html")
	})

	api.Get("/request-post", func(ctx iris.Context) {
		ctx.View("blog/request-post.html")
	})
	api.Post("/request-post", c.RequestPost)

	api.Post("/create-post", c.CreatePost)
	api.Post("/create-series", c.CreateSeries)
	api.Post("/create-comment", c.CreateComment)
	api.Post("/update-post", c.UpdatePost)
	// api.Post("/request", c.CreateRequest)
}
