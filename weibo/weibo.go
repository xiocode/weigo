/**
 * Author: xio
 * Date: 13-2-27
 * Version: 0.01
 */
package weibo

import (
	// "json"
	"fmt"
	// "log"
)

const (
	HTTP_GET    int = 0
	HTTP_POST   int = 1
	HTTP_UPLOAD int = 2
)

//METHOD_MAP = { 'GET': _HTTP_GET, 'POST': _HTTP_POST, 'UPLOAD': _HTTP_UPLOAD }

func httpCall(url string, method int, authorization string, kws ...interface{}) {
	fmt.Println(url)
	fmt.Println(method)
	fmt.Println(authorization)
	fmt.Println(kws)
}

func HttpGet(url string, authorization string, kws ...interface{}) {
	return httpCall(url, HTTP_GET, authorization, kws)
}

func HttpPost(url string, authorization string, kws ...interface{}) {
	return httpCall(url, HTTP_POST, authorization, kws)
}

func HttpUpload(url string, authorization string, kws ...interface{}) {
	return httpCall(url, HTTP_UPLOAD, authorization, kws)
}

type HttpObject struct {
	client APIClient
	method string
}

type APIClient struct {
	app_key       string
	app_secret    string
	redirect_uri  string
	response_type string
	domain        string
	version       string
	access_token  string
	expires       float32
}
