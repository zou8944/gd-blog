package sqlite

import (
	"gd-blog/gdlog"
	"gd-blog/sqlite/migrate"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"os"
	"time"
)

var ossBucket *oss.Bucket
var objectKey string
var dbFilePath string

func InitConfig() {
	endpoint := viper.GetString("aliyun.oss.endpoint")
	bucketName := viper.GetString("aliyun.oss.blog.bucket")
	accessKeyId := viper.GetString("aliyun.oss.accessKeyId")
	accessKeySecret := viper.GetString("aliyun.oss.accessKeySecret")
	objectKey = viper.GetString("aliyun.oss.blog.objectKey")
	dbFilePath = viper.GetString("database.filepath")

	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}
	ossBucket = bucket
}

func InitDB() {
	gdlog.Info("数据库文件初始化开始")

	lastModifiedInOss, existInOss := getObjectLastModifyTime(objectKey)
	lastModifiedInLocal, existInLocal := checkDbFileExist(dbFilePath)

	switch {
	case existInOss && existInLocal:
		// 本地和云端都有: 保留最新那个
		gdlog.Info("数据库文件同时存在于本地和云端...")
		if !lastModifiedInOss.After(lastModifiedInLocal) {
			gdlog.Info("本地更新，取本地...")
			break
		}
		gdlog.Info("云端更新，取云端")
		loadOssObject(dbFilePath, objectKey)
	case !existInOss && existInLocal:
		// 本地有，云端没有: 直接使用本地的
		gdlog.Info("数据库文件只存在本地，直接使用")
	case existInOss && !existInLocal:
		// 本地没有，云端有: 直接使用云端的
		gdlog.Info("数据库文件只存在云端，取云端")
		loadOssObject(dbFilePath, objectKey)
	case !existInOss && !existInLocal:
		// 本地和云端都没有: 创建新的
		gdlog.Info("本地和云端都不存在数据库文件，新建")
		_, err := os.Create(dbFilePath)
		if err != nil {
			panic(err)
		}
	}
	gdlog.Info("数据库文件初始化完成")
	migrate.AutoMigrate()
	BackupDB()
}

func loadOssObject(dbFilePath, objectKey string) {
	r, err := ossBucket.GetObject(objectKey)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	bytes, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(dbFilePath, bytes, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func checkDbFileExist(dbFilePath string) (time.Time, bool) {
	if fileInfo, err := os.Stat(dbFilePath); err != nil {
		if os.IsNotExist(err) {
			return time.Time{}, false
		}
		panic(err)
	} else {
		return fileInfo.ModTime(), true
	}
}

func getObjectLastModifyTime(objectKey string) (time.Time, bool) {
	meta, err := ossBucket.GetObjectMeta(objectKey)
	if err != nil {
		if err.(oss.ServiceError).StatusCode == 404 {
			return time.Time{}, false
		}
		panic(err)
	}
	lastModifyTimeMillis := meta.Get("Last-Modified")
	t, err := time.Parse(time.RFC1123, lastModifyTimeMillis)
	if err != nil {
		panic(err)
	}
	return t, true
}

func BackupDB() {
	gdlog.Info("数据库备份开始")
	dbFile, err := os.Open(dbFilePath)
	defer dbFile.Close()
	err = ossBucket.PutObject(objectKey, dbFile)
	if err != nil {
		panic(err)
	}
	gdlog.Info("数据库备份完成")
}
