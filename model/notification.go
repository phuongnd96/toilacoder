package model

import "time"

type Notification struct {
	TableName []byte    `json:"table_name" sql:"toilacoder.notification"`
	Id        int64     `json:"id"`
	Email     string    `json:"email" sql:",unique"`
	Series    int64     `json:"series"`
	Tag       []string  `pg:",array"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

type RequestPost struct {
	TableName []byte    `json:"table_name" sql:"toilacoder.request_post"`
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
