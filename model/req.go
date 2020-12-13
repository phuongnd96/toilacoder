package model

import (
	"time"
)

type Requestt struct {
	TableName   []byte    `json:"table_name" sql:"toilacoder.request"`
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Contact     string    `json:"contact"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateRequestt struct {
	Name        string
	Contact     string
	Description string
}
