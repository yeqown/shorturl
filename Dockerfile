#####
# build section
#####

FROM golang:1.9-alpine AS build

WORKDIR /go/src/shorturl

COPY . ./shorturl

RUN CGO_ENABLED=0 GOOS=linux ARCH=amd64 go build -o shorturl web.go

#####
# deploy section
#####

FROM alpine

WORKDIR /usr/src/shorturl

COPY --from=build /go/src/shorturl/web ./
COPY ./config.json ./

# CMD ["./web"]