package migrate

import (
	"gd-blog/gdlog"
	"gd-blog/repo/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestMigrate(t *testing.T) {
	gdlog.Info("数据库自动迁移开始")
	db, err := gorm.Open(sqlite.Open("/Users/zouguodong/Code/Personal/gd-blog/sqlite/blog.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Category{}, &model.Tag{}, &model.Visitor{})
	err = db.AutoMigrate(&model.Blog{}, &model.Comment{})
	if err != nil {
		panic(err)
	}
	gdlog.Info("数据库自动迁移完成")
}
