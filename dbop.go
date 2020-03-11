package main

import (
	"fmt"
	"sync"
	"time"
)

type dsInfo struct {
	Driver string
	Dsn    string
	Name   string
}

var dsns map[string]dsInfo
var dl sync.RWMutex

type DbInfo struct {
	Id     string
	DbUser string
	DbPass string
	DbHost string
	DbPort string
	DbName string
}

func LoadDSNs() {
	dbs := []DbInfo{
		DbInfo{
			Id:     "zhxudev",
			DbUser: "root",
			DbPass: "Connext@0101",
			DbHost: "10.128.0.180",
			DbPort: "3306",
			DbName: "cloudproject",
		},
	}
	dl.Lock()
	dsns = make(map[string]dsInfo)
	dl.Unlock()
	go func() {
		for {
			for _, di := range dbs {
				dsn := MySqlDsn(di)
				dl.Lock()
				dsns[di.Id] = dsInfo{Driver: "mysql", Dsn: dsn, Name: di.Id}
				dl.Unlock()
			}
			<-time.After(time.Minute * 2)
		}
	}()
}

func MySqlDsn(di DbInfo) string {
	dsn_tpl := "%s:%s@tcp(%s:%s)/%s"
	return fmt.Sprintf(dsn_tpl, di.DbUser, di.DbPass, di.DbHost,
		di.DbPort, di.DbName)
}