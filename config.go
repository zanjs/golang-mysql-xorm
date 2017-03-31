package main

import (
	"database/sql"

	"github.com/go-xorm/xorm"
)

// Config is ...
var Config = struct {
	APP struct {
		Port    int    `default:"12300"`
		Version string `default:"0.0.1"`
	}
	DB struct {
		UserName  string `default:"root"`
		PassWord  string `default:"root"`
		DBName    string `default:"gotest"`
		TableName string `default:"users"`
	}
	Message struct {
		Success string `default:"success"`
		Error   string `default:"error"`
	}
}{}

// DB is ...
var DB *sql.DB

// MysqlSRC is ...
var MysqlSRC = ""

// Engine is ...
var Engine *xorm.Engine
