package config

import (
	"github.com/kelseyhightower/envconfig"
)

type DBSetting struct {
	DBHost            string `envconfig:"POSTGRES_HOST" default:"database"`
	DBPort            int    `envconfig:"POSTGRES_PORT" default:"5432"`
	DBUser            string `envconfig:"POSTGRES_USER" default:"postgres"`
	DBPassword        string `envconfig:"POSTGRES_PASSWORD" default:"postgres"`
	DBName            string `envconfig:"POSTGRES_DB" default:"postgres"`
	DBSSLMode         string `envconfig:"POSTGRES_SSLMODE" default:"disable"`
}

func NewDBSetting() (DBSetting, error) {
	var dbsetting DBSetting
	err := envconfig.Process("", &dbsetting)

	return dbsetting, err
}
