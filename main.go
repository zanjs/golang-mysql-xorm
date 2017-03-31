package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/configor"
)

func init() {

	configor.Load(&Config, "config.yml")

	MysqlSRC = Config.DB.UserName + ":" + Config.DB.PassWord + "@/" + Config.DB.DBName + "?charset=utf8"

	fmt.Println(MysqlSRC)

	DB, _ = sql.Open("mysql", MysqlSRC)
	DB.SetMaxIdleConns(2000)
	DB.SetMaxIdleConns(1000)
	DB.Ping()

	Engine, _ = xorm.NewEngine("mysql", MysqlSRC)

	DataInit()

}

func main() {
	Usage()
	// pay()
	// 获取表数据信息
	// orr, _ := Engine.DBMetas()

	// fmt.Println(reflect.TypeOf(orr))

	StartHTTPServer()
}

// Student is ...
// type Student struct {
// 	Name    string
// 	Age     int
// 	Guake   bool
// 	Classes []string
// 	Price   float32
// }

// func pay() {

// 	st := &Student{
// 		"Xiao Ming",
// 		16,
// 		true,
// 		[]string{"Math", "English", "Chinese"},
// 		9.99,
// 	}

// 	s2 := new(Student)

// 	s2.Age = 16

// 	fmt.Println(s2)
// 	fmt.Println(st)
// }
