package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	hr "github.com/julienschmidt/httprouter"
	"github.com/yeqown/shorturl"
)

func main() {
	// connect to Mysql and Redis
	shorturl.ConnectDB("yeqown:yeqown@/shorturl")
	shorturl.ConnectRedis(&shorturl.RedisConfig{
		Addr:     "localhost:6379",
		Pwd:      "",
		DB:       2,
		PoolSize: 5,
	})

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

// Route ...
type Route struct {
	Path   string
	Method string
	Fn     hr.Handle
}

// Routes ...
type Routes []Route

// RegRouter ...
func RegRouter() *hr.Router {
	routes := Routes{
		{"/long/:url", http.MethodGet, HandleShortURL},
		{"/short", http.MethodGet, HandleLongURL},
	}

	router := hr.New()

	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.Fn)
	}
	return router
}

// HandleShortURL ...
func HandleShortURL(w http.ResponseWriter, req *http.Request, ps hr.Params) {
	surl := ps.ByName("url")
	originURL, err := shorturl.Parse(surl)
	if err != nil {
		responseJSON(w, &JSON{1, err.Error(), nil})
		return
	}
	if originURL == "" {
		responseJSON(w, JSON{1, "Not Existed Url", nil})
		return
	}
	log.Printf("Got request from shorturl, target originURL is: [%s]\n", originURL)
	// 301 redirect
	w.WriteHeader(http.StatusPermanentRedirect)
	http.Redirect(w, req, originURL, http.StatusPermanentRedirect)
	// http.RedirectHandler(url, code)
	// responseJson(w, &Json{0, "OK", originURL})
	// return
}

// HandleLongURL ...
func HandleLongURL(w http.ResponseWriter, req *http.Request, _ hr.Params) {
	req.ParseForm()
	longURL := req.FormValue("url")
	if longURL == "" {
		responseJSON(w, &JSON{1, "Need Param `longURL`", nil})
		return
	}
	su, err := shorturl.Shorten(longURL)
	if err != nil {
		responseJSON(w, &JSON{1, err.Error(), nil})
		return
	}
	responseJSON(w, JSON{0, "OK", su})

}

// JSON ...
type JSON struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// responseJSON ...
func responseJSON(w http.ResponseWriter, v interface{}) {
	msg, _ := json.Marshal(v)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(msg)
}
