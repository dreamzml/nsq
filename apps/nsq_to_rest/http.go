package main

import (
	//"bytes"
	"fmt"
	"time"
	"io"
	"net/http"
	"github.com/nsqio/nsq/internal/version"
	"github.com/nsqio/nsq/internal/http_api"
)

var httpclient *http.Client
var userAgent string

func init() {
	
	httpConnectTimeout :=  2*time.Second
	httpRequestTimeout := 20*time.Second

	userAgent = fmt.Sprintf("nsq_to_http v%s", version.Binary)
	httpclient = &http.Client{Transport: http_api.NewDeadlineTransport(httpConnectTimeout, httpRequestTimeout), Timeout: httpRequestTimeout}
}

func HTTPGet(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	// for key, val := range validCustomHeaders {
	// 	req.Header.Set(key, val)
	// }
	return httpclient.Do(req)
}

func HTTPPost(endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return nil, err
	}
	contentType := "application/x-www-form-urlencoded"
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", contentType)
	// for key, val := range validCustomHeaders {
	// 	req.Header.Set(key, val)
	// }
	return httpclient.Do(req)
}
