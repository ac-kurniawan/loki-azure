package common

import (
	"github.com/gojek/heimdall/v7/httpclient"
	"time"
)

func GetHttpClient() *httpclient.Client {
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(10000 * time.Millisecond),
	)
	return client
}
