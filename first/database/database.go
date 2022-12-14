package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func Connect() error {
	var err error

	dsn := "host=localhost user=root password=1234 dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}
	fmt.Println("연결 성공")
	return nil

}
