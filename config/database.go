package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBSetting struct {
	DBHost            string `envconfig:"POSTGRES_HOST" default:"localhost"`
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

func DBConnect() *gorm.DB {
	dbsetting, err := NewDBSetting()
	if err != nil {
		fmt.Println("Not create DCSetting!")
	}

	dsn := []string{
		"host=" + dbsetting.DBHost,
		"port=" + strconv.Itoa(dbsetting.DBPort),
		"user=" + dbsetting.DBUser,
		"dbname=" + dbsetting.DBName,
		"password=" + dbsetting.DBPassword,
		"sslmode=" + dbsetting.DBSSLMode,
	}
	db, err := gorm.Open(postgres.Open(strings.Join(dsn, " ")), &gorm.Config{})

	if err != nil {
		fmt.Println("Db connect failed!")
	}
	return db
}
