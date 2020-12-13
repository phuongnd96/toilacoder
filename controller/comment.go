package controller

import (
	"log"
	"regexp"
	"time"

	"toilacoder/constant"
	"toilacoder/model"

	"github.com/kataras/iris"
)

func (c *Controller) CreateComment(ctx iris.Context) {

	var createComment model.CreateComment
	err := ctx.ReadJSON(&createComment)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Định dạng sai")
		return
	}

	if createComment.Username == "" || createComment.Content == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Tên và nội dung không được để trống")
		return
	}

	var post model.PostDetail
	_, err = c.DB.Query(&post, `
		SELECT post.id, post.slug, post.description, post.content, post.view_count, post.thumbnail
		FROM toilacoder.post post
		WHERE slug = ?
	`, createComment.Slug)
	if err != nil {
		log.Println(err)
		ctx.JSON("Bài viết không tồn tại")
		return
	}

	var comment model.Comment
	comment.CreatedAt = time.Now()
	comment.ParentId = post.Id

	// Remove special characters: allow word, number, space
	reg, err := regexp.Compile(constant.REGEXPCharacters)
	if err != nil {
		ctx.JSON("Không regexp được description")
		log.Println(err)
		return
	}
	comment.Username = reg.ReplaceAllString(createComment.Username, " ")
	comment.Content = reg.ReplaceAllString(createComment.Content, " ")

	if (createComment.Email != "") {
		// Regex Email
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if (!re.MatchString(createComment.Email)) {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON("Email không hợp lệ")
			return
		}
	}
	comment.Email = createComment.Email

	err = c.DB.Insert(&comment)
	if err != nil {
		log.Println("Đường dẫn đã tồn tại", err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Đường dẫn đã tồn tại")
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(post.Slug)
}
