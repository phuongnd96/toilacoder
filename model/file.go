package model

import "time"

type Image struct {
	TableName []byte    `json:"table_name" sql:"toilacoder.image"`
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Filename  string    `json:"filename" sql:",unique"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
	Author    int64     `json:"author"`
	Tag       []string  `pg:",array"`
}

type UploadImage struct {
	Title    string
	Filename string
	Url      string
	Replace  int32
}

type DeleteImage struct {
	Id int64
}

type ListImage struct {
	Id       int64
	Title    string
	Filename string
	Source   string
}
