package main

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_parseQueryString(t *testing.T) {
	r1 := httptest.NewRequest("GET", "http://aabb.com:8888/?backend=target.com", nil)
	t.Log(r1.Host, r1.URL.Hostname())
	backend, err := parseRequest("backend", r1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("url.Host, url.Hostname()", backend.Host, backend.Hostname())

	s1 := "www.baidu.com"
	u1, err := url.Parse(s1)
	t.Log(u1, err, u1.Scheme, u1.Host, u1.Hostname())
	if u1.Scheme == "" {
		u1.Scheme = "http"
	}
	u2, err := url.Parse(u1.String())
	t.Log("u2", err, u2.Host, u2.Hostname())
}
