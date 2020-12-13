package model

type Author struct {
	TableName []byte `json:"table_name" sql:"toilacoder.author"`
	Id        int64  `json:"id"`
	Name      string `json:"name"`
}
