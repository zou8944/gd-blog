package sqlite

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path"
	"runtime"
	"testing"
)

func projectPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(path.Dir(path.Dir(filename)))
}

func InitTest() error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path.Join(projectPath(), "configs/dev.yaml"))
	err := viper.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "读取测试配置文件失败")
	}
	ak := viper.GetString("aliyun.oss.accessKeySecret")
	if len(ak) < 5 {
		return errors.New("请为测试配置文件设置access key secret，否则无法开始测试")
	}
	// 对文件路径修改，以免影响到正常的文件
	viper.Set("database.filepath", path.Join(projectPath(), "sqlite/test.db"))
	viper.Set("aliyun.oss.blog.objectKey", "blog_test.db")

	return Init()
}

func TestInitDB(t *testing.T) {
	if err := InitTest(); err != nil {
		t.Error(err)
		return
	}
	objectKey := viper.GetString("aliyun.oss.blog.objectKey")
	dbFilePath := viper.GetString("database.filepath")
	// 本地文件删除，云端也删除
	_ = os.Remove(dbFilePath)
	_ = ossBucket.DeleteObject(objectKey)
	// 跑一次，则本地有了，云端也有了

	if err := InitDB(); err != nil {
		t.Error(err)
		return
	}

	if _, err := os.Stat(dbFilePath); err != nil {
		t.Error(err)
		return
	}

	if _, err := ossBucket.GetObjectMeta(objectKey); err != nil {
		t.Error(err)
		return
	}
	// 清场
	_ = os.Remove(dbFilePath)
	_ = ossBucket.DeleteObject(objectKey)
}
