package main

import (
	"errors"
	"fmt"
	"time"
)

// User is
type User struct {
	Id      int64
	Name    string
	Age     int
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

// DataInit is ...
func DataInit() {

	isUse, _ := Engine.IsTableExist(Config.DB.TableName)

	if isUse {
		fmt.Println("表已存在")
	} else {

		fmt.Println("创建表...")

		u := new(User)

		err := Engine.CreateTables(u)

		fmt.Println(err)

		tableOp()
	}

}

// 插入测试数据
func tableOp() {

	user := User{Name: "anl", Age: 18, Created: time.Now()}

	cnt, err := Engine.Table(Config.DB.TableName).Insert(&user)

	CheckErr(err)

	if cnt != 1 {
		err = errors.New("insert not returned 1")
		fmt.Println(err)

	}

}
