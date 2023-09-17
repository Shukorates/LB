package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type server struct {
	Name         string
	URL          string
	ReverseProxy *httputil.ReverseProxy
	health       bool
}

func newServer(name, urlStr string) *server {

	u, _ := url.Parse(urlStr)
	rp := httputil.NewSingleHostReverseProxy(u)
	return &server{
		Name:         name,
		URL:          urlStr,
		ReverseProxy: rp,
		health:       true,
	}
}

func (s *server) checkHealth() bool {
	res, err := http.Head(s.URL)

	if err != nil {
		s.health = false
		return s.health
	}
	if res.StatusCode != http.StatusOK {
		s.health = false
		return s.health
	}
	s.health = true
	return s.health
}
