package sqlite

import (
	"gd-blog/configs"
	"gd-blog/internal/gdlog"
	"gd-blog/internal/sqlite/migrate"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gopkg.in/errgo.v2/errors"
	"io"
	"io/ioutil"
	"os"
	"time"
)

var ossBucket *oss.Bucket

func Init() error {
	endpoint := configs.Oss.Endpoint
	bucketName := configs.Oss.Blog.Bucket
	accessKeyId := configs.Oss.AccessKeyId
	accessKeySecret := configs.Oss.AccessKeySecret

	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return errors.Wrap(err)
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return errors.Wrap(err)
	}
	ossBucket = bucket
	return InitDB()
}

func InitDB() error {
	objectKey := configs.Oss.Blog.ObjectKey
	dbFilePath := configs.Database.FilePath
	gdlog.Info("数据库文件初始化开始")

	lastModifiedInOss, existInOss, err := getObjectLastModifyTime(objectKey)
	lastModifiedInLocal, existInLocal, err := checkDbFileExist(dbFilePath)
	if err != nil {
		return errors.Wrap(err)
	}

	switch {
	case existInOss && existInLocal:
		// 本地和云端都有: 保留最新那个
		gdlog.Info("数据库文件同时存在于本地和云端...")
		if !lastModifiedInOss.After(lastModifiedInLocal) {
			gdlog.Info("本地更新，取本地...")
			break
		}
		gdlog.Info("云端更新，取云端")
		err = loadOssObject(dbFilePath, objectKey)
	case !existInOss && existInLocal:
		// 本地有，云端没有: 直接使用本地的
		gdlog.Info("数据库文件只存在本地，直接使用")
	case existInOss && !existInLocal:
		// 本地没有，云端有: 直接使用云端的
		gdlog.Info("数据库文件只存在云端，取云端")
		err = loadOssObject(dbFilePath, objectKey)
	case !existInOss && !existInLocal:
		// 本地和云端都没有: 创建新的
		gdlog.Info("本地和云端都不存在数据库文件，新建")
		_, err := os.Create(dbFilePath)
		if err != nil {
			return errors.Wrap(err)
		}
	}
	if err != nil {
		return errors.Wrap(err)
	}
	gdlog.Info("数据库文件初始化完成")
	migrate.AutoMigrate()
	return BackupDB()
}

func loadOssObject(dbFilePath, objectKey string) error {
	r, err := ossBucket.GetObject(objectKey)
	if err != nil {
		return errors.Wrap(err)
	}
	defer func(reader io.ReadCloser) { _ = reader.Close() }(r)
	bytes, err := io.ReadAll(r)
	if err != nil {
		return errors.Wrap(err)
	}
	return ioutil.WriteFile(dbFilePath, bytes, os.ModePerm)
}

func checkDbFileExist(dbFilePath string) (time.Time, bool, error) {
	if fileInfo, err := os.Stat(dbFilePath); err != nil {
		if os.IsNotExist(err) {
			return time.Time{}, false, nil
		}
		return time.Time{}, false, errors.Wrap(err)
	} else {
		return fileInfo.ModTime(), true, nil
	}
}

func getObjectLastModifyTime(objectKey string) (time.Time, bool, error) {
	meta, err := ossBucket.GetObjectMeta(objectKey)
	if err != nil {
		if err.(oss.ServiceError).StatusCode == 404 {
			return time.Time{}, false, nil
		}
		return time.Time{}, false, errors.Wrap(err)
	}
	lastModifyTimeMillis := meta.Get("Last-Modified")
	t, err := time.Parse(time.RFC1123, lastModifyTimeMillis)
	if err != nil {
		return time.Time{}, false, errors.Wrap(err)
	}
	return t, true, nil
}

func BackupDB() error {
	objectKey := configs.Oss.Blog.ObjectKey
	dbFilePath := configs.Database.FilePath
	gdlog.Info("数据库备份开始")
	dbFile, err := os.Open(dbFilePath)
	defer func(dbFile *os.File) { _ = dbFile.Close() }(dbFile)
	err = ossBucket.PutObject(objectKey, dbFile)
	if err != nil {
		return errors.Wrap(err)
	}
	gdlog.Info("数据库备份完成")
	return nil
}
