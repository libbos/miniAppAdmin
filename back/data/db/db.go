package db

import (
	"fmt"
	"OrderMeal/data/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"OrderMeal/conf"
)

var engine *xorm.Engine

func GetOrm() *xorm.Engine {
	var err error
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", conf.Cfg("db", "user"), conf.Cfg("db", "pass"), conf.Cfg("db", "host"), conf.Cfg("db", "port"), conf.Cfg("db", "name"), conf.Cfg("db", "char"))
	engine, err = xorm.NewEngine("mysql", sqlInfo)
	engine.ShowSQL(true)
	if err != nil {
		log.Error("SQL ERROR: ", err)
		panic(err)
	}
	err = engine.Ping()
	if err != nil {
		log.Error("SQL PING ERROR: ", err)
		panic(err)
	}
	fmt.Println("数据库链接成功")
	return engine
}
