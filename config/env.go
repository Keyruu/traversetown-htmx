package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	Migrate       bool   `mapstructure:"MIGRATE" required:"true"`
	Environment   string `mapstructure:"ENVIRONMENT" required:"true"`
	BaseUrl       string `mapstructure:"BASE_URL" required:"true"`
	S3AccessKey   string `mapstructure:"S3_ACCESS_KEY" required:"true"`
	S3SecretKey   string `mapstructure:"S3_SECRET_KEY" required:"true"`
	S3Region      string `mapstructure:"S3_REGION" required:"true"`
	S3Bucket      string `mapstructure:"S3_BUCKET" required:"true"`
	S3Endpoint    string `mapstructure:"S3_ENDPOINT" required:"true"`
	S3PathStyle   bool   `mapstructure:"S3_PATH_STYLE" required:"true"`
	BackupsCron   string `mapstructure:"BACKUPS_CRON" required:"true"`
	BackupsKeep   int    `mapstructure:"BACKUPS_KEEP" required:"true"`
	BackupsBucket string `mapstructure:"BACKUPS_BUCKET" required:"true"`
	ImgproxyUrl   string `mapstructure:"IMGPROXY_URL" required:"true"`
	ImgproxyKey   string `mapstructure:"IMGPROXY_KEY" required:"true"`
	ImgproxySalt  string `mapstructure:"IMGPROXY_SALT" required:"true"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
