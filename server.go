package main

import (
	"gd-blog/config"
	"gd-blog/facade/controller"
	"gd-blog/facade/route"
	"gd-blog/facade/service"
	"gd-blog/ioc"
	"gd-blog/repo"
	mysqlite "gd-blog/sqlite"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func initComponents() error {
	// 初始化顺序：数据库 -> 几个RepoImpl -> domainService
	dbLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		LogLevel: logger.Info,
	})
	db, err := gorm.Open(sqlite.Open(config.Database.FilePath), &gorm.Config{Logger: dbLogger})
	if err != nil {
		return err
	}
	blogRepo := repo.NewBlogRepo(db)
	commentRepo := repo.NewCommentRepo(db)
	blogService := service.NewBlogService(blogRepo)
	commentService := service.NewCommentService(commentRepo)
	blogController := controller.NewBlogController(blogService)
	commentController := controller.NewCommentController(commentService)
	err = ioc.PutIn("blogController", blogController)
	err = ioc.PutIn("commentController", commentController)
	return err
}

func initAll() (*gin.Engine, error) {
	// 初始化顺序：配置、数据库package、数据库文件、系统组件、路由
	if err := config.Init(); err != nil {
		return nil, err
	}
	if err := mysqlite.Init(); err != nil {
		return nil, err
	}
	if err := mysqlite.InitDB(); err != nil {
		return nil, err
	}
	if err := initComponents(); err != nil {
		return nil, err
	}
	if r, err := route.Init(); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func main() {
	r, err := initAll()
	if err != nil {
		panic(err)
	}
	log.Fatalln("HTTP服务启动失败", r.Run(":15000"))
}
