package migrate

import (
	"gd-blog/gdlog"
	model2 "gd-blog/repo/model"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func AutoMigrate() {
	gdlog.Info("数据库自动迁移开始")
	dbFilePath := viper.GetString("database.filepath")
	db, err := gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model2.Category{}, &model2.Tag{}, &model2.Visitor{})
	err = db.AutoMigrate(&model2.Blog{}, &model2.Comment{})
	if err != nil {
		panic(err)
	}
	gdlog.Info("数据库自动迁移完成")
}
