package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// StartHTTPServer is router
func StartHTTPServer() {

	router := httprouter.New()

	// http.Handle("/", http.FileServer(http.Dir("./frontend")))
	router.GET("/pool", Pool)
	router.GET("/api/users", Users)
	router.POST("/api/users", UserCreate)
	router.GET("/api/users/:id", UserOne)
	router.PUT("/api/users/:id", UserUpdate)
	router.POST("/api/users/:id", UserDel)

	router.NotFound = http.FileServer(http.Dir("frontend"))
	err := http.ListenAndServe(":"+strconv.Itoa(Config.APP.Port), router)
	CheckErr(err)
}
