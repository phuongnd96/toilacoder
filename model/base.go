package model

import (
	"fmt"

	"toilacoder/config"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// UserClaim chứa thông tin lưu trong JWT
func ConnectDb(user string, password string, database string, address string) (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     address,
	})

	return db
}

func MigrationDb(db *pg.DB, config config.Config) error {
	/*------------- Blog --------------*/
	var image Image
	err := createTable(&image, "toilacoder", "image", db)
	if err != nil {
		return err
	}

	var post Post
	err = createTable(&post, "toilacoder", "post", db)
	if err != nil {
		return err
	}

	var comment Comment
	err = createTable(&comment, "toilacoder", "comment", db)
	if err != nil {
		return err
	}

	var notification Notification
	err = createTable(&notification, "toilacoder", "notification", db)
	if err != nil {
		return err
	}

	var reqPost RequestPost
	err = createTable(&reqPost, "toilacoder", "request_post", db)
	if err != nil {
		return err
	}

	var series Series
	err = createTable(&series, "toilacoder", "series", db)
	if err != nil {
		return err
	}

	var req Requestt
	err = createTable(&req, "toilacoder", "request", db)
	if err != nil {
		return err
	}
	var author Author
	err = createTable(&author, "toilacoder", "author", db)
	if err != nil {
		return err
	}

	var tag Tag
	err = createTable(&tag, "toilacoder", "tag", db)
	if err != nil {
		return err
	}

	return nil
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}

// LogQueryToConsole Log câu lệnh query
func LogQueryToConsole(db *pg.DB) {
	db.AddQueryHook(dbLogger{})
}

func createTable(model interface{}, schema, tableName string, db *pg.DB) error {
	exist, err := TableIsExists(schema, tableName, db)
	if err != nil {
		return err
	}
	if !exist {
		err = db.CreateTable(model, &orm.CreateTableOptions{
			Temp:          false,
			FKConstraints: true,
			IfNotExists:   true,
		})
		if err != nil {
			return err
		}

	}

	return err
}

// TableIsExists kiểm tra bảng đã tồn tại
func TableIsExists(schema, tableName string, db *pg.DB) (bool, error) {
	var exist bool
	_, err := db.Query(&exist, `
		SELECT EXISTS (
			SELECT 1
			FROM   information_schema.tables 
			WHERE  table_schema = ?
			AND    table_name = ?
			)`, schema, tableName)
	if err != nil {
		return exist, err
	}
	return exist, err
}
