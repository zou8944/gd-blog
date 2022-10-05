package migrate

import (
	"gd-blog/configs"
	"gd-blog/internal/gdlog"
	"gd-blog/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func AutoMigrate() {
	gdlog.Info("数据库自动迁移开始")
	db, err := gorm.Open(sqlite.Open(configs.Database.FilePath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Category{}, &models.Tag{})
	err = db.AutoMigrate(&models.Blog{}, &models.Comment{})
	if err != nil {
		panic(err)
	}
	gdlog.Info("数据库自动迁移完成")
}
