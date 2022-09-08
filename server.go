package main

import (
	"gd-blog/facade/controller"
	"gd-blog/facade/route"
	"gd-blog/facade/service"
	"gd-blog/ioc"
	"gd-blog/repo"
	mysqlite "gd-blog/sqlite"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config/dev.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	mysqlite.InitConfig()
}

func initComponents() {
	// 初始化顺序：数据库 -> 几个RepoImpl -> domainService
	dbLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		LogLevel: logger.Info,
	})
	dbFilePath := viper.GetString("database.filepath")
	db, err := gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{Logger: dbLogger})
	if err != nil {
		panic(err)
	}
	blogRepo := repo.NewBlogRepo(db)
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)
	err = ioc.PutIn("blogController", blogController)
	if err != nil {
		panic(err)
	}
}

func initAll() *gin.Engine {
	initConfig()
	mysqlite.InitDB()
	initComponents()
	r, err := route.Init()
	if err != nil {
		panic(err)
	}
	return r
}

func main() {
	r := initAll()
	log.Fatalln("HTTP服务启动失败", r.Run(":15000"))
}
