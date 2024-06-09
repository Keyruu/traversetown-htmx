package migrations

import (
	"github.com/keyruu/traversetown-htmx/config"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		env := config.NewEnv()
		dao := daos.New(db)

		settings, _ := dao.FindSettings()
		settings.Meta.AppName = "traversetown"
		settings.Logs.MaxDays = 5
		settings.S3.Enabled = true
		settings.S3.AccessKey = env.S3AccessKey
		settings.S3.Secret = env.S3SecretKey
		settings.S3.Region = env.S3Region
		settings.S3.Bucket = env.S3Bucket
		settings.S3.Endpoint = env.S3Endpoint
		settings.S3.ForcePathStyle = env.S3PathStyle

		settings.Backups.S3.Enabled = true
		settings.Backups.S3.AccessKey = env.BackupsS3AccessKey
		settings.Backups.S3.Secret = env.BackupsS3SecretKey
		settings.Backups.S3.Region = env.BackupsS3Region
		settings.Backups.S3.Bucket = env.BackupsBucket
		settings.Backups.S3.Endpoint = env.BackupsS3Endpoint
		settings.Backups.S3.ForcePathStyle = env.BackupsS3PathStyle
		settings.Backups.Cron = env.BackupsCron
		settings.Backups.CronMaxKeep = env.BackupsKeep

		return dao.SaveSettings(settings)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		settings, _ := dao.FindSettings()
		settings.S3.Enabled = false
		settings.S3.AccessKey = ""
		settings.S3.Secret = ""
		settings.S3.Region = ""
		settings.S3.Bucket = ""
		settings.S3.Endpoint = ""
		settings.S3.ForcePathStyle = false

		settings.Backups.S3.Enabled = false
		settings.Backups.S3.AccessKey = ""
		settings.Backups.S3.Secret = ""
		settings.Backups.S3.Region = ""
		settings.Backups.S3.Bucket = ""
		settings.Backups.S3.Endpoint = ""
		settings.Backups.S3.ForcePathStyle = false
		settings.Backups.Cron = ""
		settings.Backups.CronMaxKeep = 0

		return dao.SaveSettings(settings)
	})
}
