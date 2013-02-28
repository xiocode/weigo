/**
 * Author: xio
 * Date: 13-2-27
 * Version: 0.01
 */
package weibo

import (
	// "json"
	"fmt"
	"net/url"
	// "strconv"
	"strings"
	"time"
	// "log"
)

const (
	HTTP_GET    int = 0
	HTTP_POST   int = 1
	HTTP_UPLOAD int = 2
)

/************************************************************
*
* Http Request *
*
*************************************************************/

func httpCall(url string, method int, authorization string, kws map[string]string) bool {
	var params string
	var boundary string
	var http_url string
	var http_body string

	if method == HTTP_UPLOAD {
		url = strings.Replace(url, "https://api.", "https://upload.api.", 1)
		params, boundary = encodeMultipart()
	} else {
		params = encodeParams(kws)
	}

	if method == HTTP_GET {
		http_url = fmt.Sprintf("%s?%s", url, params)
		http_body = ""
	} else {
		http_url = url
		http_body = params
	}

	fmt.Println(http_url, http_body)

	if authorization != "" {
		fmt.Println(authorization)
	}
	if boundary != "" {
		fmt.Println(boundary)
	}
	return true
}

/************************************************************
*
* Encode Params *
*
*************************************************************/
func encodeParams(kws map[string]string) (params string) {
	if len(kws) > 0 {
		values := url.Values{}
		for key, value := range kws {
			values.Add(key, value)
		}
		params = values.Encode()
	}
	return
}

/************************************************************
*
* Encode For Upload *
*
*************************************************************/
func encodeMultipart() (params string, boundary string) {
	return "", ""
}

/************************************************************
*
* HTTPObject Struct *
*
*************************************************************/
type HttpObject struct {
	client *APIClient
	method int
}

/************************************************************
*
* Call httpcall function, make a request *
*
*************************************************************/
func (http *HttpObject) Call(uri string, kws map[string]string) {
	fmt.Println(http.client, http.method)
	var url = fmt.Sprintf("%s%s.json", http.client.api_url, uri)
	httpCall(url, http.method, "authorize", kws)
}

/************************************************************
*
* APIClient Struct *
*
*************************************************************/
type APIClient struct {
	app_key       string
	app_secret    string
	redirect_uri  string
	response_type string
	domain        string
	auth_url      string
	api_url       string
	version       string
	access_token  string
	expires       int64
	Get           *HttpObject
	Post          *HttpObject
	Upload        *HttpObject
}

/************************************************************
*
* Check Is Expires *
*
*************************************************************/
func (api *APIClient) is_expires() bool {
	return api.access_token == "" || api.expires < time.Now().Unix()
}

/************************************************************
*
* Create New API Client Instance & Init *
*
*************************************************************/
func NewAPIClient(app_key, app_secret, redirect_uri string) *APIClient {
	var api = &APIClient{
		app_key:       app_key,
		app_secret:    app_secret,
		redirect_uri:  redirect_uri,
		response_type: "code",
		domain:        "api.weibo.com",
		version:       "2",
	}
	api.auth_url = fmt.Sprintf("https://%s/oauth2/", api.domain)
	api.api_url = fmt.Sprintf("https://%s/%s/", api.domain, api.version)
	api.Get = &HttpObject{client: api, method: HTTP_GET}
	api.Post = &HttpObject{client: api, method: HTTP_POST}
	api.Upload = &HttpObject{client: api, method: HTTP_UPLOAD}
	return api
}

type APIError struct {
	when       time.Time
	error_code int
	message    string
}

func (err *APIError) Error() string {
	return fmt.Sprintf("APIError When: %v Message: %v Code: %v", err.when, err.message, err.error_code)
}
