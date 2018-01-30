package shorturl

import (
	hr "github.com/julienschmidt/httprouter"

	"encoding/json"
	"log"
	"net/http"
)

type Route struct {
	Path   string
	Method string
	Fn     hr.Handle
}

type Routes []Route

func RegRouter() *hr.Router {
	routes := Routes{
		{"/long/:short_url", http.MethodGet, HandleShortUrl},
		{"/short", http.MethodGet, HandleLongUrl},
	}

	router := hr.New()

	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.Fn)
	}
	return router
}

func HandleShortUrl(w http.ResponseWriter, req *http.Request, ps hr.Params) {
	shorturl := ps.ByName("short_url")
	longurl, err := Parse(shorturl)
	if err != nil {
		responseJson(w, &Json{1, err.Error(), nil})
		return
	}
	if longurl == "" {
		responseJson(w, Json{1, "Not Existed Url", nil})
		return
	}
	log.Printf("Got request from shorturl, target longurl is: [%s]\n", longurl)
	// 301 redirect
	w.WriteHeader(http.StatusPermanentRedirect)
	http.Redirect(w, req, longurl, http.StatusPermanentRedirect)
	// http.RedirectHandler(url, code)
	// responseJson(w, &Json{0, "OK", longurl})
	// return
}

func HandleLongUrl(w http.ResponseWriter, req *http.Request, _ hr.Params) {
	req.ParseForm()
	longurl := req.FormValue("longurl")
	if longurl == "" {
		responseJson(w, &Json{1, "Need Param `longurl`", nil})
		return
	}
	su, err := Shorten(longurl)
	if err != nil {
		responseJson(w, &Json{1, err.Error(), nil})
		return
	}
	responseJson(w, Json{0, "OK", su})

}

type Json struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func responseJson(w http.ResponseWriter, v interface{}) {
	msg, _ := json.Marshal(v)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(msg)
}
