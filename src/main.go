package main

import (
	"database/sql"
	"gd-blog/src/domain/service"
	"gd-blog/src/ioc"
	"gd-blog/src/repoimpl"
	"gd-blog/src/route"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func initComponents() error {
	// 初始化顺序：数据库 -> 几个RepoImpl -> domainService
	db, err := sql.Open("sqlite3", "file:sqlite/blog.db")
	if err != nil {
		return err
	}
	blogRepoImpl := repoimpl.NewBlogRepoImpl(db)
	commentRepoImpl := repoimpl.NewCommentRepoImpl(db)
	userRepoImpl := repoimpl.NewUserRepoImpl(db)
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
		Addr: ":9090",
	}
	log.Println("监听中 :9090")
	log.Fatalln("HTTP服务启动失败", s.ListenAndServe())
}
