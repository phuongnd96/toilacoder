package main

import (
	"fmt"
	"os"

	"toilacoder/config"
	"toilacoder/controller"
	"toilacoder/helper"
	"toilacoder/model"
	"toilacoder/router"

	"github.com/go-pg/pg"
	"github.com/kataras/iris"
)

func main() {

	config := config.SetupConfig()

	// Khởi tạo controller
	c := controller.NewController()
	c.Config = config

	// Kết nối CSDL
	dbConfig := config.Database
	db := model.ConnectDb(dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Address)
	defer db.Close()
	c.DB = db
	setupDatabase(db, config)

	//Khởi tạo app iris
	app := iris.New()
	app.Logger().SetLevel("debug")

	// Đăng ký phân quyền
	// app.Use(c.WrapRouter)

	// Serve static files (css)
	app.StaticWeb("/resources", "./resources")
	app.StaticWeb("/images", "./images")

	// // Đăng ký thư mục chứa HTML
	tmpl := iris.HTML("./view", ".html")

	// // Đăng ký thư mục chứa layout
	tmpl.Layout("layout/layout.html").Reload(true)

	tmpl.AddFunc("listMenu", helper.ListMenu)
	tmpl.AddFunc("inc", func(i int) int {
		return i + 1
	})
	app.RegisterView(tmpl)

	// Đăng ký các group router
	router.BasicRoutes(c, app)
	router.FileRoutes(c, app)
	router.PostRoutes(c, app)

	app.Favicon("./resources/image/favicon.png")
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func setupDatabase(db *pg.DB, config config.Config) {
	argsWithProg := os.Args
	if len(argsWithProg) > 1 && os.Args[1] == "release" {
	} else {
		model.LogQueryToConsole(db)
	}

	err := model.MigrationDb(db, config)
	if err != nil {
		fmt.Printf("hehe")
		panic(err)
	}
}
