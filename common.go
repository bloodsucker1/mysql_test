package main

import (
	"database/sql"
	"fmt"
	"github.com/ngaut/log"
)

type DBConfig struct {
	user     string
	password string
	host     string
	port     int
}

func (dbConfig *DBConfig) checkAndReconnection(db *sql.DB) (*sql.DB, error) {
	if db == nil {
		var err error
		db, err = dbConfig.connect()
		return db, err
	}
	if db.Ping() != nil {
		log.Debugf("reconnect user:%s addr:%s:%d", dbConfig.user, dbConfig.host, dbConfig.port)
		return dbConfig.connect()
	}
	return db, nil
}

func (dbConfig *DBConfig) connect() (*sql.DB, error) {
	return sql.Open("mysql", dbConfig.user+":"+dbConfig.password+"@tcp("+fmt.Sprintf("%s:%d", dbConfig.host, dbConfig.port)+")/?charset=utf8&autocommit=true&timeout=1s")
}