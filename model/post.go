package model

import (
	"html/template"
	"time"
)

type Post struct {
	TableName   []byte    `json:"table_name" sql:"toilacoder.post"`
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug" sql:",unique"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Thumbnail   string    `json:"thumbnail"`
	ViewCount   int64     `json:"view_count" sql:",notnull"`
	Status      int32     `json:"status" sql:",notnull"`
	CreatedAt   time.Time `json:"created_at"`
	AuthorId    int64     `json:"author_id"`
	Tag         []string  `pg:",array"`
	Category    string    `json:"category"`
	Series      int64     `json:"series"`
}

type GetPost struct {
	Id              int64
	Slug            string
	Title           string
	Thumbnail       string
	Description     string
	CreatedAtString string
	Author          string
	AuthorId        int64
	Tag             []string `pg:",array"`
	Category        string
}

type GetTags struct {
	Id   int
	Name string
	Slug string
}

type PostDetail struct {
	Id          int64
	Slug        string
	Title       string
	Thumbnail   string
	Content     template.HTML
	Link        string
	Description string
	CreatedAt   string
	Author      string
	AuthorId    int64
	Tag         []string `pg:",array"`
	ViewCount   int64
	Type        string
	SeriesId    int64
}

type CreatePost struct {
	Content     string
	Title       string
	Slug        string
	Tag         string
	Description string
	Thumbnail   string
	SeriesId    int64
}

type UpdatePost struct {
	Id          int64
	Content     string
	Title       string
	Tag         string
	Description string
	Thumbnail   string
	Slug        string
	SeriesId    int64
}

type CreateConfirmTheNotification struct {
	Email string
}

type RequestPostReq struct {
	Email string
	Content string
}
