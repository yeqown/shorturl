# Shorturl library

ShortUrl System Coding


### Quick start

```sh
go run main.go --dsn 'test:test@tcp(127.0.0.1:3306)/shorten_url'
2021/01/26 13:57:03 listening on :8080
```

```sh
curl http://localhost:8080/api/shorten?l=http://www.baidu.com
```