package db

import (
	"log"
	"time"

	"ecnu.space/tmp-loser/conf"
	"github.com/go-xorm/xorm"
)

var (
	dsn          = conf.GetAppConfig().MysqlConfig.Dsn
	maxIdleConns = conf.GetAppConfig().MysqlConfig.MaxIdleConns
	maxOpenConns = conf.GetAppConfig().MysqlConfig.MaxOpenConns
)

func initMySQL() {
	Engine = NewMysql()
}

func NewMysql() *xorm.Engine {
	db, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatalf("NewMysql NewEngine Error: %v, dsn: %s", err, dsn)
	}
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.ShowSQL()
	db.DatabaseTZ, err = time.LoadLocation("UTC")
	if err != nil {
		log.Fatalf("NewMysql failed to set application timezone %v, dsn: %s", err, dsn)
	}
	db.TZLocation, err = time.LoadLocation("Local")
	if err != nil {
		log.Fatalf("NewMysql failed to set application timezone %v, dsn: %s", err, dsn)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("NewMysql cannot connect to database %v, dsn: %s", err, dsn)
	}
	return db
}
