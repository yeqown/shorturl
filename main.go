package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/yeqown/shorturl/pkg/base62"
	"github.com/yeqown/shorturl/pkg/orm"

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

	*domain = *domain + "/r?s="
	impl, err := construct(*domain, *dsn)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/shorten", impl.Shorten)
	http.HandleFunc("/r", impl.Parse)

	log.Printf("listening on :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

type shortURLImpl struct {
	domain string

	orm *orm.ShortORM
}

func construct(domain, dsn string) (*shortURLImpl, error) {
	v, err := orm.NewORM(dsn, "mysql")
	if err != nil {
		return nil, err
	}
	return &shortURLImpl{
		domain: domain,
		orm:    v,
	}, nil
}

func (s shortURLImpl) assembleURL(encoded string) string {
	return s.domain + encoded
}

// Shorten ...
func (s shortURLImpl) Shorten(w http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	source := req.FormValue("l")
	if len(source) == 0 {
		_ = s.responseJSON(w, errors.New("empty source URL (as l)"))
		return
	}

	m := orm.ShortURLDO{
		Source: source,
	}
	if err := s.orm.Create(&m); err != nil {
		_ = s.responseJSON(w, err)
		return
	}

	shorted := s.assembleURL(base62.Encode(m.ID))
	go func() {
		(&m).Shorted = shorted
		if err := s.orm.Update(&m); err != nil {
			fmt.Printf("ERR: could not update shorted: %+v", m)
		}
	}()

	_ = s.responseJSON(w, shorted)
}

// Parse short_url convert 2 id
func (s shortURLImpl) Parse(w http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	short := req.FormValue("s")
	if len(short) == 0 {
		_ = s.responseJSON(w, errors.New("empty shorted URL (as s)"))
		return
	}

	id := base62.Decode(short)
	out := orm.ShortURLDO{
		ID: id,
	}

	if err := s.orm.Query(&out); err != nil {
		_ = s.responseJSON(w, err)
		return
	}

	_ = s.responseJSON(w, out.Source)
}

// responseJSON ...
func (s shortURLImpl) responseJSON(w http.ResponseWriter, v interface{}) error {
	if err, ok := v.(error); ok {
		v = withError(err)
	} else {
		v = withData(v)
	}

	msg, _ := json.Marshal(v)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write(msg)
	return err
}

// responseWrapper ...
type responseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func withError(err error) responseWrapper {
	return responseWrapper{
		Code:    -1,
		Message: "Failed: " + err.Error(),
		Data:    nil,
	}
}

func withData(v interface{}) responseWrapper {
	return responseWrapper{
		Code:    0,
		Message: "Success",
		Data:    v,
	}
}
