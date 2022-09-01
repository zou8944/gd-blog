package main

import (
	"gd-blog/src/domain/service"
	"gd-blog/src/ioc"
	"gd-blog/src/persistence"
	"gd-blog/src/route"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"log"
	"net/http"
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
	blogRepoImpl := persistence.NewBlogRepoImpl(db)
	commentRepoImpl := persistence.NewCommentRepoImpl(db)
	userRepoImpl := persistence.NewUserRepoImpl(db)
	blogDomainService := service.NewBlogDomainService(blogRepoImpl, commentRepoImpl, userRepoImpl)
	blogHandler := route.NewBlogHandler(blogDomainService)
	err = ioc.PutIn("blogDomainService", blogDomainService)
	err = ioc.PutIn("blogHandler", blogHandler)
	return err
}

func main() {
	err := initComponents()
	if err != nil {
		log.Fatalln("IOC初始化失败", err)
	}
	err = route.InitRoutes()
	if err != nil {
		log.Fatalln("路由初始化失败", err)
	}
	s := &http.Server{
		Addr: ":15000",
	}
	log.Println("监听中 :15000")
	log.Fatalln("HTTP服务启动失败", s.ListenAndServe())
}
