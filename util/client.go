package util

import (
	"net/http"
	"time"
)

//HTTPClient is a default HTTP Client
var HTTPClient = http.Client{
	Timeout: time.Duration(time.Second * 10),
}

//IsConnected is used to check internet available
func IsConnected() bool {
	if _, err := HTTPClient.Get("https://www.google.com"); err != nil {
		return false
	}
	return true
}
