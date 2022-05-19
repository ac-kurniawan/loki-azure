package common

import (
	"github.com/gojek/heimdall/v7/httpclient"
	"time"
)

func GetHttpClient() *httpclient.Client {
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(1000 * time.Millisecond),
	)
	return client
}
