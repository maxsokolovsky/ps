package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

const domain = "localhost"

var baseUrl = fmt.Sprintf("https://%s%s", domain, *addr)

func getClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
