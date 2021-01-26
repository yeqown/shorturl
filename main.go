package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/yeqown/shorturl/internal"

	_ "github.com/go-sql-driver/mysql"
)

var (
	domain = flag.String("domain", "http://localhost:8080",
		"specify domain as shorter URL host ")
	dsn  = flag.String("dsn", "USER:PASSWORD@/DBNAME", "addr to connect to MySQL")
	port = flag.Int("addr", 8080, "addr to listen on")
)

func main() {
	flag.Parse()

	impl, err := internal.NewShorten(*domain, *dsn)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/shorten", impl.Shorten)
	http.HandleFunc("/r", impl.Parse)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
