package config

import (
	"be21/users"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	var connectionString = "host=localhost user=postgres password=TEAMSECRETGG dbname=db_library port=6666 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("terjadi sebuah kesalahan", err.Error())
		return nil
	}
	return db
}

func Migrate(connection *gorm.DB) error {
	err := connection.AutoMigrate(&users.User{})
	return err
}
