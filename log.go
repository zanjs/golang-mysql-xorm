package main

import (
	"os"

	"github.com/go-xorm/xorm"
)

// Log is ...
func Log() {

	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	Engine.SetLogger(xorm.NewSimpleLogger(f))

}
