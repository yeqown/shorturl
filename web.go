package main

import (
	"log"
	"net/http"
	"time"

	. "shorturl/_shorturl"
)

func main() {
	LoadConfig("./_shorturl/config.json")
	ins := GetInstance()
	// connect to Mysql and Redis
	ConnectDB(ins.MySql)
	ConnectRedis(ins.Redis)
	router := RegRouter()
	// add server timeout setting
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":9090",
		Handler:      router,
	}
	log.Println("server listen on: ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
