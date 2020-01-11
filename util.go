package main

import (
	"log"
	"net/http"
	"net/url"
	"time"

	cache "github.com/patrickmn/go-cache"
)

var internalCache *cache.Cache

func init() {
	internalCache = cache.New(24*time.Hour, 10*time.Minute)
}

func parseRequest(queryStringKey string, r *http.Request) (backend *url.URL, err error) {
	backend, err = url.Parse(r.FormValue(queryStringKey))
	return
}

// 1. lookup in cache using request host name as key, if NOT found
// 2. parse request url(query string) for next request, AND cache it.
func findBackend(queryStringKey string, r *http.Request) (backend *url.URL, err error) {
	cacheKey := r.Host
	val, found := internalCache.Get(cacheKey)
	if found {
		backend, err = url.Parse(val.(string))
		return
	}
	backend, err = parseRequest(queryStringKey, r)
	if err != nil {
		return
	}
	if backend.Scheme == "" {
		backend.Scheme = "http"
		backend, _ = url.Parse(backend.String())
	}
	log.Println("set cache", cacheKey, backend.String())
	internalCache.Set(cacheKey, backend.String(), 0)
	return
}
