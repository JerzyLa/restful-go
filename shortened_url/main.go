package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/lytics/base62"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var ShortenedUrlMap = make(map[string]string)

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func createURL(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := buf.String()
	if !isValidUrl(body) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	shortenedUrl := base62.StdEncoding.EncodeToString([]byte(body))
	ShortenedUrlMap[shortenedUrl] = body

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, shortenedUrl)
}

func redirectURL(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	shortenedUrl := params.ByName("url")

	if originalUrl, ok := ShortenedUrlMap[shortenedUrl]; !ok {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusMovedPermanently)
		_, _ = fmt.Fprintf(w, originalUrl)
	}
}

func main() {
	router := httprouter.New()
	router.POST("/api/v1/new", createURL)
	router.GET("/api/v1/:url", redirectURL)
	log.Fatal(http.ListenAndServe(":8000", router))
}
