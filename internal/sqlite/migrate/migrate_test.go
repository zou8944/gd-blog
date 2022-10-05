package migrate

import (
	"gd-blog/internal/gdlog"
	"gd-blog/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestMigrate(t *testing.T) {
	gdlog.Info("数据库自动迁移开始")
	db, err := gorm.Open(sqlite.Open("/Users/zouguodong/Code/Personal/gd-blog/blog.db"), &gorm.Config{})
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
