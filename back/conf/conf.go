package conf

import (
	"gopkg.in/ini.v1"
	"log"
)

func Cfg(sec, key string) string {
	var cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		log.Fatal(err)
	}
	return cfg.Section(sec).Key(key).String()
}
