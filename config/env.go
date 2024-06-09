package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	Migrate            bool   `mapstructure:"MIGRATE" required:"true"`
	Environment        string `mapstructure:"ENVIRONMENT" required:"true"`
	BaseUrl            string `mapstructure:"BASE_URL" required:"true"`
	S3AccessKey        string `mapstructure:"S3_ACCESS_KEY" required:"true"`
	S3SecretKey        string `mapstructure:"S3_SECRET_KEY" required:"true"`
	S3Region           string `mapstructure:"S3_REGION" required:"true"`
	S3Bucket           string `mapstructure:"S3_BUCKET" required:"true"`
	S3Endpoint         string `mapstructure:"S3_ENDPOINT" required:"true"`
	S3PathStyle        bool   `mapstructure:"S3_PATH_STYLE" required:"true"`
	BackupsCron        string `mapstructure:"BACKUPS_CRON" required:"true"`
	BackupsKeep        int    `mapstructure:"BACKUPS_KEEP" required:"true"`
	BackupsS3Bucket    string `mapstructure:"BACKUPS_S3_BUCKET" required:"true"`
	BackupsS3AccessKey string `mapstructure:"BACKUPS_S3_ACCESS_KEY" required:"true"`
	BackupsS3SecretKey string `mapstructure:"BACKUPS_S3_SECRET_KEY" required:"true"`
	BackupsS3Region    string `mapstructure:"BACKUPS_S3_REGION" required:"true"`
	BackupsS3Endpoint  string `mapstructure:"BACKUPS_S3_ENDPOINT" required:"true"`
	BackupsS3PathStyle bool   `mapstructure:"BACKUPS_S3_PATH_STYLE" required:"true"`
	ImgproxyUrl        string `mapstructure:"IMGPROXY_URL" required:"true"`
	ImgproxyKey        string `mapstructure:"IMGPROXY_KEY" required:"true"`
	ImgproxySalt       string `mapstructure:"IMGPROXY_SALT" required:"true"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Can't find the file .env : ", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Println("Environment can't be loaded: ", err)
	}

	return &env
}
