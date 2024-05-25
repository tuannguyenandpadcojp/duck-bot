package utils

import (
	"net/http"
	"time"
)

func NewHTTPClient() *http.Client {
	tp := http.DefaultTransport.(*http.Transport).Clone()
	tp.MaxIdleConnsPerHost = 10

	return &http.Client{
		Timeout:   10 * time.Second,
		Transport: tp,
	}
}
