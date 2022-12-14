package configs

import (
	"github.com/spf13/viper"
	"gopkg.in/errgo.v2/errors"
	"os"
)

type oss struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Blog            blog
}

type blog struct {
	Bucket    string
	ObjectKey string
}

type database struct {
	FilePath string
}

var Oss oss
var Database database

func Init() error {
	var configFilePath string
	switch viper.GetString("env") {
	case "test":
		configFilePath = "configs/test.yaml"
	case "prod":
		configFilePath = "configs/prod.yaml"
	default:
		configFilePath = "configs/dev.yaml"
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err)
	}

	type configs struct {
		Oss      oss
		Database database
	}
	var c configs

	if err := viper.Unmarshal(&c); err != nil {
		return errors.Wrap(err)
	}
	Oss = c.Oss
	Database = c.Database
	pathFromEnv := os.Getenv("SQLITE_FILE")
	if pathFromEnv != "" {
		Database.FilePath = pathFromEnv
	}

	return nil
}
