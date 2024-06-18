package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type Env struct {
	Migrate             bool   `env:"MIGRATE,required"`
	Environment         string `env:"ENVIRONMENT,required"`
	BaseUrl             string `env:"BASE_URL,required"`
	S3AccessKey         string `env:"S3_ACCESS_KEY,required"`
	S3SecretKey         string `env:"S3_SECRET_KEY,required"`
	S3Region            string `env:"S3_REGION,required"`
	S3Bucket            string `env:"S3_BUCKET,required"`
	S3Endpoint          string `env:"S3_ENDPOINT,required"`
	S3PathStyle         bool   `env:"S3_PATH_STYLE,required"`
	BackupsCron         string `env:"BACKUPS_CRON,required"`
	BackupsKeep         int    `env:"BACKUPS_KEEP,required"`
	BackupsS3Bucket     string `env:"BACKUPS_S3_BUCKET,required"`
	BackupsS3AccessKey  string `env:"BACKUPS_S3_ACCESS_KEY,required"`
	BackupsS3SecretKey  string `env:"BACKUPS_S3_SECRET_KEY,required"`
	BackupsS3Region     string `env:"BACKUPS_S3_REGION,required"`
	BackupsS3Endpoint   string `env:"BACKUPS_S3_ENDPOINT,required"`
	BackupsS3PathStyle  bool   `env:"BACKUPS_S3_PATH_STYLE,required"`
	ImgproxyUrl         string `env:"IMGPROXY_URL,required"`
	ImgproxyKey         string `env:"IMGPROXY_KEY,required"`
	ImgproxySalt        string `env:"IMGPROXY_SALT,required"`
	LastfmUrl           string `env:"LASTFM_URL,required"`
	LastfmApiKey        string `env:"LASTFM_API_KEY,required"`
	SpotifyClientId     string `env:"SPOTIFY_CLIENT_ID,required"`
	SpotifyClientSecret string `env:"SPOTIFY_CLIENT_SECRET,required"`
	SpotifyRefreshToken string `env:"SPOTIFY_REFRESH_TOKEN,required"`
}

func NewEnv() *Env {
	envStruct := Env{}

	err := env.Parse(&envStruct) // 👈 Parse environment variables into `Config`
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	return &envStruct
}
