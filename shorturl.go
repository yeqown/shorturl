package shorturl

import (
	"log"

	"github.com/yeqown/shorturl/base62"
)

// Domain ...
const Domain = "localhost:9090/long"

// Shorten ...
func Shorten(long string) (string, error) {
	var err error
	db, err := GetDB()
	if err != nil {
		return "", err
	}

	// get from cache first
	if id, ex := checkURLCacheExist(long); ex {
		um := URLModel{
			DB: db,
			ID: id,
		}
		if err := um.query(); err != nil {
			log.Println("Query mysql got an err: ", err)
			return "", err
		}
		if um.ShortURL != "" {
			// reset url cache key expire
			setURLCache(um.LongURL, um.ID)
			return um.ShortURL, nil
		}
		log.Printf("Get shorturl from cahe with id=[%d]: `%s` but is empty string\n", um.ID, um.ShortURL)
		delURLCache(um.LongURL)
	}

	// if not target long in cache
	um := URLModel{
		DB:      db,
		LongURL: long,
	}
	if _, err = um.insert(); err != nil {
		return "", err
	}
	base62 := base62.Encode(um.ID)
	um.ShortURL = Domain + "/" + base62
	if err := um.update(); err != nil {
		return "", err
	}
	setURLCache(um.LongURL, um.ID)
	return um.ShortURL, err
}

// Parse short_url convert 2 id
func Parse(short string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}

	id := base62.Decode(short)
	um := &URLModel{
		ID: id,
		DB: db,
	}
	if err := um.query(); err != nil {
		return "", err
	}
	return um.LongURL, nil
}
