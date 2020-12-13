package controller

import (
	"fmt"
	"log"
	"strings"
	"time"
	"toilacoder/constant"
	"toilacoder/model"

	"github.com/kataras/iris"
	"github.com/rivo/uniseg"
)

func (c *Controller) GetSeriesBySlug(ctx iris.Context) {
	slug := ctx.Params().Get("slug")

	auth, _ := constant.Sess.Start(ctx).GetBoolean("authenticated")
	if auth {
		ctx.ViewData("Logged", true)
	}

	var series model.SeriesDetail
	_, err := c.DB.Query(&series, `
		SELECT series.id, series.slug, series.description, series.content, series.thumbnail,
		series.tag, 
		series.title,
		to_char(series.created_at, 'DD/MM/YYYY') created_at
		FROM toilacoder.series series
		WHERE slug = ?
	`, slug)
	if err != nil {
		log.Println(err)
		ctx.JSON("Bài viết không tồn tại")
		return
	}

	if series.Id == 0 {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON("Not Found")
		return
	}

	if series.Description == "" {
		series.Description = series.Title
	}

	tagsString := strings.Join(series.Tag, ",")

	var metaDescription string
	metaDescription = series.Description
	// Giới hạn độ dài thẻ meta description
	nummberChars := uniseg.GraphemeClusterCount(metaDescription)
	if nummberChars > 158 {
		metaDescription = metaDescription[:158]
	}

	var postsInSeries []model.GetPost
	query := fmt.Sprintf(`
		SELECT post.id, post.slug, post.description, post.thumbnail, 
		post.title, post.author_id, post.category, post.tag,
		to_char(cast(post.created_at as date), 'DD-MM-YYYY') as created_at_string
		FROM toilacoder.post post
		WHERE post.series = ?
		ORDER BY created_at DESC 
	`)
	_, err = c.DB.Query(&postsInSeries, query, series.Id)
	if err != nil {
		log.Println(err)
		ctx.JSON("Lỗi phía server")
		return
	}

	ctx.ViewData("ogURL", constant.DOMAIN+"/series/"+slug)
	ctx.ViewData("title", series.Title)
	ctx.ViewData("ogTitle", series.Title)
	ctx.ViewData("ogImage", series.Thumbnail)
	ctx.ViewData("ogDescription", metaDescription)
	ctx.ViewData("metaDescription", metaDescription)
	ctx.ViewData("metaKeyword", ", "+tagsString)

	ctx.ViewData("series", series)
	ctx.ViewData("postsInSeries", postsInSeries)
	ctx.View("series/series.html")
}

func (c *Controller) CreateSeries(ctx iris.Context) {
	auth, _ := constant.Sess.Start(ctx).GetBoolean("authenticated")
	if !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	var createSeries model.CreateSeries
	err := ctx.ReadJSON(&createSeries)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Định dạng sai")
		return
	}

	if createSeries.Title == "" || createSeries.Slug == "" || createSeries.Description == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Tiêu đề, Mô tả, Đường dẫn không được để trống")
		return
	}

	var series model.Series
	series.Title = createSeries.Title
	series.Slug = createSeries.Slug
	series.Description = createSeries.Description
	series.Thumbnail = createSeries.Thumbnail
	series.Content = createSeries.Content
	series.CreatedAt = time.Now()

	s := strings.Split(createSeries.Tag, ",")

	for i, x := range s {
		if (x == "") || (x == "\"\"") {
			s = remove(s, i)
		}
	}
	series.Tag = s

	err = c.DB.Insert(&series)
	if err != nil {
		log.Println("Đường dẫn đã tồn tại", err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Đường dẫn đã tồn tại")
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(series.Slug)
}
