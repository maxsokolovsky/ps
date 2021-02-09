package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

var baseUrl = fmt.Sprintf("https://%s%s", *host, *addr)

func getClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
