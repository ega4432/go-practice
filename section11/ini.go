package main

import (
	"fmt"
	"gopkg.in/go-ini/ini.v1"
)

// https://godoc.org/gopkg.in/go-ini/ini.v1

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	// config.ini を読み込む
	cfg, _ := ini.Load("section11/config.ini")
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(),
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
