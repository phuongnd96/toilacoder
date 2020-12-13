package model

type Tag struct {
	TableName []byte `json:"table_name" sql:"toilacoder.tag"`
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
}
