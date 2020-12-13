package model

import (
	"html/template"
	"time"
)

type Series struct {
	TableName     []byte    `json:"table_name" sql:"toilacoder.series"`
	Id            int64     `json:"id"`
	Title         string    `json:"title"`
	Slug          string    `json:"slug" sql:",unique"`
	Description   string    `json:"description"`
	Content       string    `json:"content"`
	Thumbnail     string    `json:"thumbnail"`
	Status        int32     `json:"status" sql:",notnull"`
	CreatedAt     time.Time `json:"created_at"`
	AuthorId      int64     `json:"author_id"`
	Tag           []string  `pg:",array"`
}

type GetSeries struct {
	Id              int64
	Slug            string
	Title           string
	Thumbnail       string
	Description     string
	CreatedAtString string
	Author          string
	AuthorId        int64
	Tag             []string `pg:",array"`
}

type SeriesDetail struct {
	Id            int64
	Slug          string
	Title         string
	Thumbnail     string
	Content       template.HTML
	Link          string
	Description   string
	CreatedAt     string
	Author        string
	AuthorId      int64
	Tag           []string `pg:",array"`
}

type CreateSeries struct {
	Content     string
	Title       string
	Slug        string
	Tag         string
	Description string
	Thumbnail   string
}