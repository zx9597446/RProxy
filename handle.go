package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type handle struct {
	queryStringKey string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("request:", r.URL.Host+r.URL.String())
	remoteHost, err := findBackend(this.queryStringKey, r)
	if err != nil || remoteHost == "" {
		log.Println("findBackend", err, "remoteHost:", remoteHost)
		return
	}
	remoteURL, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println("parse url:", err)
		return
	}
	remoteURL.Host = remoteHost
	remoteURL.Path = "/"
	if remoteURL.Scheme == "" {
		remoteURL.Scheme = "http"
		remoteURL, _ = url.Parse(remoteURL.String())
	}
	log.Println("proxy", r.URL.Host+r.URL.String(), "->", remoteURL.String())
	proxy := httputil.NewSingleHostReverseProxy(remoteURL)
	r.Host = remoteHost
	proxy.ServeHTTP(w, r)
}
