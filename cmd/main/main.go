package main

import (
	"gd-blog/configs"
	mysqlite "gd-blog/internal/sqlite"
	"gd-blog/pkg/models"
	"gd-blog/pkg/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func initAll() error {
	// 初始化顺序：配置、sqlite数据库文件、数据库
	if err := configs.Init(); err != nil {
		return err
	}
	if err := mysqlite.Init(); err != nil {
		return err
	}
	if err := models.Init(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := initAll(); err != nil {
		panic(err)
	}
	r := gin.New()
	routes.RegisterRoutes(r)
	log.Fatalln("HTTP服务启动失败", r.Run(":15000"))
}
