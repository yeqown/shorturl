package shorturl

import (
	"log"
)

const Domain = "localhost:9090/long"

func Shorten(longurl string) (string, error) {
	var err error
	db, err := GetDB()
	if err != nil {
		return "", err
	}

	// get from cache first
	if id, ex := CheckUrlCacheExist(longurl); ex {
		um := UrlModel{
			DB: db,
			Id: id,
		}
		if err := um.Query(); err != nil {
			log.Println("Query mysql got an err: ", err)
			return "", err
		}
		if um.ShortUrl != "" {
			// resert url cache key expire
			SetUrlCache(um.LongUrl, um.Id)
			return um.ShortUrl, nil
		}
		log.Printf("Get shorturl from cahe with id=[%d]: `%s` but is empty string\n", um.Id, um.ShortUrl)
		DelUrlCache(um.LongUrl)
	}

	// if not target longurl in cache
	um := UrlModel{
		DB:      db,
		LongUrl: longurl,
	}
	if _, err = um.Insert(); err != nil {
		return "", err
	}
	base62 := Encode(um.Id)
	um.ShortUrl = Domain + "/" + base62
	if err := um.Update(); err != nil {
		return "", err
	}
	SetUrlCache(um.LongUrl, um.Id)
	return um.ShortUrl, err
}

// short_url convert 2 id
func Parse(shorturl string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}

	id := Decode(shorturl)
	um := &UrlModel{
		Id: id,
		DB: db,
	}
	if err := um.Query(); err != nil {
		return "", err
	}
	return um.LongUrl, nil
}

func splitUrl() {

}
