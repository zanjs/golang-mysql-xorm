package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Pool is ...
func Pool(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users := ModelUser()

	// 类型打印
	fmt.Println(reflect.TypeOf(users))
	// map 返回json 编码
	str, err := json.Marshal(users)

	CheckErr(err)

	fmt.Fprintln(w, string(str))
}

// Users is ...
func Users(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// users := ModelUser()

	users := make([]User, 0)

	err := Engine.Find(&users)
	// users := new(User)
	// users, err := Engine.Table(Config.DB.UserName).Get(&User{Name: "anl"})
	CheckErr(err)
	fmt.Println(users)

	fmt.Fprintln(w, ReJSON(users))

}

// UserOne is ...
func UserOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// r.FormValue()

	id, _ := strconv.ParseInt(ps.ByName("id"), 10, 64)

	fmt.Println(id)

	user := &User{Id: id}
	has, _ := Engine.Get(user)
	fmt.Println(has)

	fmt.Fprintln(w, ReJSON(user))
}

// UserCreate is ...
func UserCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	name := r.FormValue("name")
	age, err := strconv.Atoi(r.FormValue("age"))

	user := User{Name: name, Age: age, Created: time.Now()}

	cnt, err := Engine.Table(Config.DB.TableName).Insert(&user)

	CheckErr(err)

	if cnt != 1 {
		err = errors.New("insert not returned 1")
		fmt.Println(err)

	}

	fmt.Fprintln(w, ReJSON(user))

}

// UserUpdate is ...
func UserUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)

	name := r.FormValue("name")
	age, err := strconv.Atoi(r.FormValue("age"))

	user := User{Name: name, Age: age, Id: id}

	cnt, err := Engine.Table(Config.DB.TableName).Id(id).Update(&user)

	CheckErr(err)

	if cnt != 1 {
		err = errors.New("Update not returned 1")
		fmt.Println(err)

	}

	fmt.Fprintln(w, ReJSON(user))
}

// UserDel is ...
func UserDel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)

	cnt, err := Engine.Table(Config.DB.TableName).Id(id).Delete(&User{})

	CheckErr(err)

	if cnt != 1 {
		err = errors.New("Del not returned 1")
		fmt.Println(err)

	}

	fmt.Fprintln(w, ReJSON(Config.Message.Success))
}
