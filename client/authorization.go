package client

import "encoding/base64"

type BasicAuth struct {
	Username string `json:"username" xml:"username"`
	Password string `json:"password" xml:"password"`
}

func HTTPAuthorizationBasic(auth BasicAuth) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth.Username+":"+auth.Password))
}
