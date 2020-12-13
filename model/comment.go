package model

import (
	"time"
)

type Comment struct {
	TableName []byte    `json:"table_name" sql:"toilacoder.comment"`
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Content   string    `json:"content"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	ParentId  int64     `json:"parent_id"`
}

type GetComment struct {
	Id        int64
	Username  string
	Email     string
	Content   string
	Avatar    string
	CreatedAt string
	ParentId  int64
}

type CreateComment struct {
	Username string
	Email    string
	Content  string
	Slug     string
}
