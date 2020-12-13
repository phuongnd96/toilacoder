package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"toilacoder/constant"
	"toilacoder/model"

	"github.com/kataras/iris"
	"github.com/rivo/uniseg"
)

func (c *Controller) GetPosts(ctx iris.Context) {
	tagParams := ctx.URLParams()
	keyword := ctx.URLParamDefault("keyword", "")
	searchKeyword := "'%" + keyword + "%'"
	currentPage, _ := strconv.Atoi(ctx.URLParamDefault("page", "1"))
	var offset = (currentPage - 1) * 10

	splitParams := strings.Split(tagParams["tag"], ",")

	var searchTags = ""
	if len(splitParams) > 0 {
		if splitParams[0] != " " {
			for _, x := range splitParams {
				searchTags = searchTags + "AND post.tag @> '{" + x + "}' "
			}
		}
	}

	var posts []model.GetPost
	query := fmt.Sprintf(`
		SELECT post.id, post.slug, post.description, post.thumbnail, 
		post.title, post.author_id, post.category, post.tag,
		to_char(cast(post.created_at as date), 'DD-MM-YYYY') as created_at_string
		FROM toilacoder.post post
		WHERE (post.title ILIKE %s OR post.description ILIKE %s )
		%s
		ORDER BY created_at DESC 
		LIMIT 10
		OFFSET ?
	`, searchKeyword, searchKeyword, searchTags)
	_, err := c.DB.Query(&posts, query, offset)
	if err != nil {
		log.Println(err)
		ctx.JSON("Lỗi phía server")
		return
	}

	for index := range posts {
		// Remove special characters: allow word, number, space
		reg, err := regexp.Compile(constant.REGEXPCharacters)
		if err != nil {
			ctx.JSON("Không regexp được description")
			log.Println(err)
			return
		}
		posts[index].Description = reg.ReplaceAllString(posts[index].Description, " ")

		// Giới hạn độ dài thẻ meta description
		nummberChars := uniseg.GraphemeClusterCount(posts[index].Description)
		if nummberChars > 125 {
			posts[index].Description = posts[index].Description[:125]
		}

	}

	var numberPage int32
	var totalPages []int32
	var totalItems int32
	query = fmt.Sprintf(`
		SELECT COUNT(*)
		FROM  toilacoder.post post 
		WHERE (post.title ILIKE %s OR post.description ILIKE %s )
		%s
	`, searchKeyword, searchKeyword, searchTags)
	_, err = c.DB.Query(&totalItems, query)
	if err != nil {
		log.Println(err)
		ctx.JSON("Lỗi phía server")
		return
	}

	if (totalItems % int32(10)) != 0 {
		numberPage = (totalItems / int32(10)) + 1
	} else {
		numberPage = (totalItems / int32(10))
	}

	var index int32 = 1
	for index <= numberPage {
		totalPages = append(totalPages, index)
		index++
	}

	var tags []string
	_, err = c.DB.Query(&tags, `
		select distinct unnest(tag) as tag
		from toilacoder.post
	`)
	if err != nil {
		log.Println(err)
		ctx.JSON("Lỗi phía server")
		return
	}

	var series []model.GetSeries
	query = fmt.Sprintf(`
		SELECT series.id, series.slug, series.description, series.thumbnail, 
		series.title, series.author_id, series.tag
		FROM toilacoder.series series
		ORDER BY created_at DESC 
	`)
	_, err = c.DB.Query(&series, query)
	if err != nil {
		log.Println(err)
		ctx.JSON("Lỗi phía server")
		return
	}

	var title = "Chia sẻ công nghệ, kiến thức, kỹ năng lập trình - Tôi là Coder"
	var description = "Tổng hợp kiến thức, chia sẻ, bug, kinh nghiệm, hướng dẫn, như thế nào về lập trình Golang, Go, Docker, Devops, Backend, Java, Javascript, Frontend"

	ctx.ViewData("ogURL", constant.DOMAIN+"/")
	ctx.ViewData("title", title)
	ctx.ViewData("ogTitle", title)
	ctx.ViewData("ogImage", constant.DOMAIN+"/resources/image/logo.png")
	ctx.ViewData("ogDescription", description)
	ctx.ViewData("metaDescription", description)

	ctx.ViewData("keyword", keyword)
	ctx.ViewData("totalPages", totalPages)
	ctx.ViewData("currentPage", currentPage)
	ctx.ViewData("listPost", posts)
	ctx.ViewData("listSeries", series)
	ctx.ViewData("listTag", tags)
	ctx.ViewData("tagChoosed", splitParams)
	ctx.View("blog/index.html")
}

func (c *Controller) GetPostBySlug(ctx iris.Context) {
	slug := ctx.Params().Get("slug")

	// auth, _ := constant.Sess.Start(ctx).GetBoolean("authenticated")
	// if auth {
	// 	ctx.ViewData("Logged", true)
	// }

	var post model.PostDetail
	_, err := c.DB.Query(&post, `
		SELECT post.id, post.slug, post.description, post.content, post.view_count, post.thumbnail,
		post.tag, post.series series_id,
		post.title,
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
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON("Not Found")
		return
	}

	_, err = c.DB.Exec(`
		UPDATE toilacoder.post
		SET view_count = ?
		WHERE slug = ? AND id = ?
	`, post.ViewCount+1, slug, post.Id)
	if err != nil {
		log.Println(err)
		ctx.JSON("Lỗi phía server")
		return
	}

	var postsInSeries []model.GetPost
	query := fmt.Sprintf(`
		SELECT post.id, post.slug, post.description, post.Title
		FROM toilacoder.post post
		WHERE series = ?
		ORDER BY created_at DESC 
	`, )
	_, err = c.DB.Query(&postsInSeries, query, post.SeriesId)
	if err != nil {
		log.Println(err)
		ctx.JSON("Lỗi phía server")
		return
	}

	if post.Description == "" {
		post.Description = post.Title
	}

	tagsString := strings.Join(post.Tag, ",")

	var metaDescription string
	metaDescription = post.Description

	// Remove special characters: allow word, number, space
	reg, err := regexp.Compile(constant.REGEXPCharacters)
	if err != nil {
		ctx.JSON("Không regexp được description")
		log.Println(err)
		return
	}
	metaDescription = reg.ReplaceAllString(metaDescription, " ")

	// Giới hạn độ dài thẻ meta description
	nummberChars := uniseg.GraphemeClusterCount(metaDescription)
	if nummberChars > 158 {
		metaDescription = metaDescription[:158]
	}

	// Tìm kiếm bài viết hiện tại thuộc series nào
	var series model.SeriesDetail
	_, err = c.DB.Query(&series, `
		SELECT series.id, series.slug, series.description, series.content, series.thumbnail, 
		series.tag, series.title,
		to_char(series.created_at, 'DD/MM/YYYY') created_at
		FROM toilacoder.series series
		WHERE series.id = ?
	`, post.SeriesId)
	if err != nil {
		log.Println(err)
		ctx.JSON("Bài viết không tồn tại")
		return
	}

	if series.Id != 0 {
		ctx.ViewData("postIsInSeries", series)
	}

	// Get comments of post
	var comments []model.GetComment
	query = fmt.Sprintf(`
		SELECT comment.id, comment.username, comment.email, comment.content, comment.avatar,
		comment.parent_id, to_char(comment.created_at, 'HH24:MI DD/MM/YYYY') created_at
		FROM toilacoder.comment comment
		WHERE comment.parent_id = ?
	`)
	_, err = c.DB.Query(&comments, query, post.Id)
	if err != nil {
		log.Println(err)
		ctx.JSON("Lỗi phía server")
		return
	}

	ctx.ViewData("ogURL", constant.DOMAIN+"/"+slug)
	ctx.ViewData("title", post.Title)
	ctx.ViewData("ogTitle", post.Title)
	ctx.ViewData("ogImage", post.Thumbnail)
	ctx.ViewData("ogDescription", metaDescription)
	ctx.ViewData("metaDescription", metaDescription)
	ctx.ViewData("metaKeyword", ", "+tagsString)

	ctx.ViewData("post", post)
	ctx.ViewData("commentsOfPost", comments)
	ctx.ViewData("postsInSeries", postsInSeries)
	ctx.View("blog/post.html")
}

func (c *Controller) CreatePost(ctx iris.Context) {
	// auth, _ := constant.Sess.Start(ctx).GetBoolean("authenticated")
	// if !auth {
	// 	ctx.StatusCode(iris.StatusForbidden)
	// 	return
	// }
	var createPost model.CreatePost
	err := ctx.ReadJSON(&createPost)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Định dạng sai")
		return
	}

	if createPost.Title == "" || createPost.Content == "" || createPost.Slug == "" || createPost.Description == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Tiêu đề, Mô tả, Đường dẫn, Nội dung không được để trống")
		return
	}

	var post model.Post
	post.Title = createPost.Title
	post.Slug = createPost.Slug
	post.Description = createPost.Description
	post.Thumbnail = createPost.Thumbnail
	post.Content = createPost.Content
	post.CreatedAt = time.Now()
	if createPost.SeriesId != 0 {
		post.Series = createPost.SeriesId
	}

	s := strings.Split(createPost.Tag, ",")

	for i, x := range s {
		if (x == "") || (x == "\"\"") {
			s = remove(s, i)
		}
	}
	post.Tag = s

	err = c.DB.Insert(&post)
	if err != nil {
		log.Println("Đường dẫn đã tồn tại", err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Đường dẫn đã tồn tại")
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(post.Slug)
}

func (c *Controller) UpdatePost(ctx iris.Context) {
	// auth, _ := constant.Sess.Start(ctx).GetBoolean("authenticated")
	// if !auth {
	// 	ctx.JSON("Forbidden")
	// 	ctx.StatusCode(iris.StatusForbidden)
	// 	return
	// }

	var req model.UpdatePost
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Lỗi cập nhật")
		return
	}

	var id int64
	_, err = c.DB.Query(&id, `
	select id from toilacoder.post where id = ?
	`, req.Id)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Lỗi cập nhật")
		return
	}

	if id == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Bài viết không tồn tại")
		return
	}

	// Update dữ liệu
	var post model.Post
	post.Id = id
	post.Title = req.Title
	post.Slug = req.Slug
	post.Thumbnail = req.Thumbnail
	post.Description = req.Description
	post.Content = req.Content
	if req.SeriesId != 0 {
		post.Series = req.SeriesId
	}

	s := strings.Split(req.Tag, ",")
	for i, x := range s {
		if (x == "") || (x == "\"\"") {
			s = remove(s, i)
		}
	}
	post.Tag = s

	_, err = c.DB.Model(&post).Column("title", "slug", "description", "content", "thumbnail", "tag", "series").WherePK().Update()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Tiêu đề có thể đã tồn tại")
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Cập nhật thành công")
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func (c *Controller) ConfirmTheNotification(ctx iris.Context) {
	var createConfirm model.CreateConfirmTheNotification
	err := ctx.ReadJSON(&createConfirm)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Định dạng sai")
		return
	}

	if createConfirm.Email == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Email không được để trống")
		return
	}

	// Regex Email
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(createConfirm.Email) {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Email không hợp lệ")
		return
	}

	var noti model.Notification
	noti.Email = createConfirm.Email
	noti.CreatedAt = time.Now()

	err = c.DB.Insert(&noti)
	if err != nil {
		log.Println("Email đã đăng ký trước đó", err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Email đã đăng ký trước đó")
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Đăng ký thành công")
}

func (c *Controller) RequestPost(ctx iris.Context) {
	var reqPost model.RequestPostReq
	err := ctx.ReadJSON(&reqPost)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Định dạng sai")
		return
	}

	if reqPost.Content == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Nội dung không được để trống")
		return
	}

	var reqP model.RequestPost
	reqP.Email = reqPost.Email
	reqP.Content = reqPost.Content
	reqP.CreatedAt = time.Now()

	err = c.DB.Insert(&reqP)
	if err != nil {
		log.Println("Không thể yêu cầu lúc này", err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON("Không thể yêu cầu lúc này")
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Cám ơn bạn! Mình sẽ cố gắng thực hiện nội dung sớm nhất có thể")
}
