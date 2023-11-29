package clients

import (
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var Jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary

// GetHttpClient returns a new http.Client with the following settings:
// - Timeout: 10 seconds
// - MaxIdleConns: 100
// - MaxConnsPerHost: 100
// - MaxIdleConnsPerHost: 100
func NewHttpClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: t,
	}
}
