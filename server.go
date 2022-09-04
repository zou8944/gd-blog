package main

import (
	"gd-blog/facade/controller"
	route2 "gd-blog/facade/route"
	"gd-blog/facade/service"
	"gd-blog/ioc"
	repo "gd-blog/repo"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"log"
	"os"
)

func initComponents() error {
	// 初始化顺序：数据库 -> 几个RepoImpl -> domainService
	dbLogger := logger2.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger2.Config{
		LogLevel: logger2.Info,
	})
	db, err := gorm.Open(sqlite.Open("sqlite/blog.db"), &gorm.Config{Logger: dbLogger})
	if err != nil {
		return err
	}
	blogRepo := repo.NewBlogRepo(db)
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)
	err = ioc.PutIn("blogController", blogController)
	return err
}

func main() {
	err := initComponents()
	if err != nil {
		log.Fatalln("IOC初始化失败", err)
	}
	r, err := route2.Init()
	if err != nil {
		log.Fatalln("路由初始化失败", err)
	}
	log.Fatalln("HTTP服务启动失败", r.Run(":15000"))
}
