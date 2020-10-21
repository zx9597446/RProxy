package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	cache "github.com/patrickmn/go-cache"
)

var internalCache *cache.Cache

func init() {
	//internalCache = cache.New(24*time.Hour, 24*time.Hour)
	internalCache = cache.New(10*time.Minute, 10*time.Minute)
}

func parseRequest(queryStringKey string, r *http.Request) (backendHost string, err error) {
	queryVal := r.FormValue(queryStringKey)
	if queryVal != "" {
		backend, err := url.Parse(queryVal)
		if err == nil {
			backendHost = backend.Host
		}
	} else {
		err = fmt.Errorf("can't find %s in query", queryStringKey)
	}
	return
}

// 1. lookup in cache using request host name as key, if NOT found
// 2. parse request url(query string) for next request, AND cache it.
func findBackend(queryStringKey string, r *http.Request) (backendHost string, err error) {
	cacheKey := r.Host
	val, found := internalCache.Get(cacheKey)
	if found {
		backendHost = val.(string)
		return
	}
	backendHost, err = parseRequest(queryStringKey, r)
	if err != nil {
		return
	}
	log.Println("set cache", cacheKey, "->", backendHost)
	internalCache.Set(cacheKey, backendHost, 0)
	return
}
