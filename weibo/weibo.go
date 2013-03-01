/**
 * Author: xio
 * Date: 13-2-27
 * Version: 0.01
 */
package weibo

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	// "strconv"
	"io"
	"os"
	"strings"
	// "sync"
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

func httpCall(the_url string, method int, authorization string, kws map[string]string) bool {
	var params string
	var boundary string
	var http_url string
	var http_body string

	if method == HTTP_UPLOAD {
		the_url = strings.Replace(the_url, "https://api.", "https://upload.api.", 1)
		params, boundary = encodeMultipart()
	} else {
		params = encodeParams(kws)
	}

	if method == HTTP_GET {
		http_url = fmt.Sprintf("%s?%s", the_url, params)
		http_body = ""
	} else {
		http_url = the_url
		http_body = params
	}

	fmt.Println(http_url, http_body)

	var HTTP_METHOD string

	switch method {
	case HTTP_GET:
		HTTP_METHOD = "GET"
	case HTTP_POST:
		HTTP_METHOD = "POST"
	case HTTP_UPLOAD:
		HTTP_METHOD = "POST"
	default:
		HTTP_METHOD = "GET"
	}

	url, err := url.Parse(http_url)
	checkError(err)

	client := new(http.Client)                                                               // New Http Client
	request, err := http.NewRequest(HTTP_METHOD, url.String(), strings.NewReader(http_body)) //Make New Request
	request.Header.Add("Accept-Encoding", "gzip")
	if authorization != "" {
		request.Header.Add("Authorization", fmt.Sprintf("OAuth2 %s", authorization))
	}
	if boundary != "" {
		request.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", boundary))
	}

	response, err := client.Do(request) // Do Request
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	checkError(err)
	defer response.Body.Close()

	body := read_body(response)
	fmt.Println(body)
	parse_json(body)
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
* Read Response Body *
*
*************************************************************/
func read_body(response *http.Response) (body string) {
	var reader io.ReadCloser
	buffer := make([]byte, 1024) //Buffer
	using_gzip := response.Header.Get("Content-Encoding")
	switch using_gzip {
	case "gzip":
		reader, err := gzip.NewReader(response.Body)
		checkError(err)
		defer reader.Close()
	default:
		reader = response.Body
	}

	for {
		chunk, err := reader.Read(buffer[0:])
		if err != nil && err != io.EOF {
			os.Exit(0)
		}
		if chunk == 0 { // End Of Body
			break
		}
		body += string(buffer[0:chunk])
	}
	return body
}

/************************************************************
*
*Parse Json To JSON Struct *
*
*************************************************************/
func parse_json(body string) (result interface{}) {
	// jsonDict := new(map[string]interface{})
	// data, err := json.NewDecoder(strings.NewReader(body))
	// checkError(err)
	data_bytes := []byte(body)
	if err := json.Unmarshal(body, &result); err == io.EOF {
		break
	} else if err != nil {
		checkError(err)
	}
	fmt.Println(result)
	return result
}

/************************************************************
*
* JSON *
*
*************************************************************/

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
	if http.client.is_expires() {
		panic(&APIError{when: time.Now(), error_code: "21327", message: "expired_token"})
	}
	httpCall(url, http.method, http.client.access_token, kws)
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
// func NewAPIClient(app_key, app_secret, redirect_uri string) *APIClient {
// 	var api = &APIClient{
// 		app_key:       app_key,
// 		app_secret:    app_secret,
// 		redirect_uri:  redirect_uri,
// 		response_type: "code",
// 		domain:        "api.weibo.com",
// 		version:       "2",
// 	}
// 	api.auth_url = fmt.Sprintf("https://%s/oauth2/", api.domain)
// 	api.api_url = fmt.Sprintf("https://%s/%s/", api.domain, api.version)
// 	api.Get = &HttpObject{client: api, method: HTTP_GET}
// 	api.Post = &HttpObject{client: api, method: HTTP_POST}
// 	api.Upload = &HttpObject{client: api, method: HTTP_UPLOAD}
// 	return api
// }

func NewAPIClient(app_key, app_secret, redirect_uri string) *APIClient {
	var api = &APIClient{
		app_key:       app_key,
		app_secret:    app_secret,
		redirect_uri:  redirect_uri,
		response_type: "code",
		domain:        "api.pgysocial.com/apiproxy/weibo",
		version:       "2",
	}
	api.auth_url = fmt.Sprintf("http://%s/oauth2/", api.domain)
	api.api_url = fmt.Sprintf("http://%s/%s/", api.domain, api.version)
	api.Get = &HttpObject{client: api, method: HTTP_GET}
	api.Post = &HttpObject{client: api, method: HTTP_POST}
	api.Upload = &HttpObject{client: api, method: HTTP_UPLOAD}
	return api
}

func (api *APIClient) SetAccessToken(access_token string, expires int64) *APIClient {
	api.access_token = access_token
	api.expires = expires
	return api
}

type APIError struct {
	when       time.Time
	error_code string
	message    string
}

func (err *APIError) Error() string {
	return fmt.Sprintf("APIError When: %v Message: %v Code: %v", err.when, err.message, err.error_code)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
