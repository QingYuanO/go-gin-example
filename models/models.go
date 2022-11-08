package models

import (
	"fmt"
	"go-gin-example/pkg/setting"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	ID         int `grom:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                       error
		dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")

	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10s", user, password, host, dbName)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // DSN data source name
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix:   tablePrefix,
		SingularTable: true,
	},
		Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Println(err)
	}
	sqlDB, _ := db.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

}

// func CloseDB() {
// 	defer db.Close()
// }
