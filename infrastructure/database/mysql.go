package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	writeConnection *gorm.DB
	initOnce        sync.Once
)

func MysqlConnection() (db *gorm.DB) {
	initOnce.Do(func() {
		writeConnection = newMysqlConnection()
	})
	return writeConnection
}

func newMysqlConnection() (db *gorm.DB) {
	fmt.Println(connectionString())
	db, err := gorm.Open("mysql", connectionString())
	if err != nil {
		panic(err.Error())
	}
	// d := config.GetEnvValue().Database
	// db.LogMode(d.LogMode)
	db.DB()
	db.DB().SetConnMaxLifetime(time.Hour * 24)
	// db.DB().SetMaxIdleConns(d.MaxIdleConnections)
	// db.DB().SetMaxOpenConns(d.MaxOpenConnections)
	return db
}

func connectionString() string {

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		"smashop",
		"smashop",
		"127.0.0.1",
		3306,
		"itemcloud",
	)
}
