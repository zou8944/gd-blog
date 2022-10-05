package import_export

import (
	"gopkg.in/errgo.v2/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestImport(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("/Users/zouguodong/Code/Personal/gd-blog/blog.db"), &gorm.Config{
		Logger: logger2.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger2.Config{
			SlowThreshold: time.Second + 1,
			Colorful:      true,
			LogLevel:      logger2.Error,
		}),
	})

	err = filepath.Walk("/Users/zouguodong/Code/Other/blog/source/_posts/", func(path string, info fs.FileInfo, err error) error {
		println(path)
		if strings.HasSuffix(path, ".md") {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			return Import(db, file)
		}
		return nil
	})

	if err != nil {
		t.Error(errors.Wrap(err))
	}
}
