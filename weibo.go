/**
 * Author: xio
 * Date: 13-2-27
 * Version: 0.01
 */
package weigo

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	// "reflect"
	"strings"
	// "sync"
	"log"
	"strconv"
	"time"
)

const (
	HTTP_GET    int = 0
	HTTP_POST   int = 1
	HTTP_UPLOAD int = 2
)

func httpCall(the_url string, method int, authorization string, params map[string]interface{}) (result map[string]interface{}) {
	var url_params string
	var multipart_data *bytes.Buffer //Upload Image
	var http_url string
	var http_body string
	var content_type string
	var multipart_content_type string
	var request *http.Request
	var HTTP_METHOD string
	var err error

	if method == HTTP_UPLOAD {
		the_url = strings.Replace(the_url, "https://api.", "https://upload.api.", 1)
		multipart_content_type, multipart_data = encodeMultipart(params)
	} else {
		url_params = encodeParams(params)
	}

	if method == HTTP_GET {
		http_url = fmt.Sprintf("%s?%s", the_url, url_params)
		http_body = ""
	} else {
		http_url = the_url
		http_body = url_params
	}

	switch method {
	case HTTP_GET:
		HTTP_METHOD = "GET"
	case HTTP_POST:
		HTTP_METHOD = "POST"
		content_type = "application/x-www-form-urlencoded"
	case HTTP_UPLOAD:
		HTTP_METHOD = "POST"
		content_type = multipart_content_type
	}

	client := new(http.Client)
	if method == HTTP_UPLOAD {
		request, err = http.NewRequest(HTTP_METHOD, http_url, multipart_data) //Upload
	} else {
		request, err = http.NewRequest(HTTP_METHOD, http_url, strings.NewReader(http_body)) //GET & POST
	}

	request.Header.Add("Accept-Encoding", "gzip")
	if method != HTTP_GET {
		request.Header.Add("Content-Type", content_type)
	}

	if authorization != "" {
		request.Header.Add("Authorization", fmt.Sprintf("OAuth2 %s", authorization))
	}

	response, err := client.Do(request) // Do Request
	checkError(err)
	defer response.Body.Close()
	// if response.Status != "200 OK" {
	// 	panic(response.Status)
	// }
	body := read_body(response)
	result = parse_json(body)

	if error_code, ok := result["error_code"].(int); ok {
		panic(&APIError{when: time.Now(), error_code: strconv.Itoa(error_code), message: result["error"].(string)})
	}

	return result
}

func encodeParams(params map[string]interface{}) (result string) {
	if len(params) > 0 {
		values := url.Values{}
		for key, value := range params {
			switch value.(type) {
			case string:
				values.Add(key, value.(string))
			case int:
				values.Add(key, strconv.Itoa(value.(int)))
			}
		}
		result = values.Encode()
	}
	return
}

func encodeMultipart(params map[string]interface{}) (multipartContentType string, multipartData *bytes.Buffer) {
	if len(params) > 0 {
		multipartData = new(bytes.Buffer)
		bufferWriter := multipart.NewWriter(multipartData) // type *bytes.Buffer
		for key, value := range params {
			switch value.(type) {
			case string:
				bufferWriter.WriteField(key, value.(string))
			case int:
				bufferWriter.WriteField(key, string(value.(int)))
			case *os.File:
				picdata, err := bufferWriter.CreateFormFile(key, value.(*os.File).Name())
				multipartContentType = bufferWriter.FormDataContentType()
				defer bufferWriter.Close()
				if err != nil {
					checkError(err)
				}
				io.Copy(picdata, value.(*os.File))
			}
		}
	}
	return multipartContentType, multipartData
}

func read_body(response *http.Response) (body string) {
	var reader io.ReadCloser
	var err error
	var contents []byte

	using_gzip := response.Header.Get("Content-Encoding")
	switch using_gzip {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		checkError(err)
		defer reader.Close()
	default:
		reader = response.Body
	}
	contents, err = ioutil.ReadAll(reader)
	if err != nil {
		checkError(err)
	}
	body = string(contents)
	return body
}

func parse_json(body string) (result map[string]interface{}) {
	var data interface{}

	body_bytes := []byte(body)
	err := json.Unmarshal(body_bytes, &data)
	if err != nil {
		checkError(err)
	}
	result = data.(map[string]interface{})
	fmt.Println(result)
	return result
}

type HttpObject struct {
	client *APIClient
	method int
}

func (http *HttpObject) Call(uri string, params map[string]interface{}) (result map[string]interface{}) {
	// fmt.Println(http.client, http.method)
	var url = fmt.Sprintf("%s%s.json", http.client.api_url, uri)
	if http.client.is_expires() {
		panic(&APIError{when: time.Now(), error_code: "21327", message: "expired_token"})
	}
	result = httpCall(url, http.method, http.client.access_token, params)
	return
}

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

func (api *APIClient) is_expires() bool {
	return api.access_token == "" || api.expires < time.Now().Unix()
}

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

func (api *APIClient) SetAccessToken(access_token string, expires int64) *APIClient {
	api.access_token = access_token
	api.expires = expires
	return api
}

func (api *APIClient) GetAuthorizeUrl(redirect_uri string, params map[string]interface{}) string {
	var redirect string
	var response_type string
	var ok bool

	if redirect_uri != "" {
		redirect = redirect_uri
	} else {
		redirect = api.redirect_uri
	}

	if redirect == "" {
		panic(&APIError{when: time.Now(), error_code: "21305", message: "Parameter absent: redirect_uri"})
	}

	if response_type, ok = params["response_type"].(string); !ok {
		response_type = "code"
	}

	var url_params = map[string]interface{}{
		"client_id":     api.app_key,
		"response_type": response_type,
		"redirect_uri":  redirect,
	}

	for key, value := range params {
		url_params[key] = value
	}

	return fmt.Sprintf("%s%s?%s", api.auth_url, "authorize",
		encodeParams(url_params))
}

func (api *APIClient) RequestAccessToken(code, redirect_uri string) (result map[string]interface{}) {
	var redirect string
	var the_url string
	the_url = fmt.Sprintf("%s%s", api.auth_url, "access_token")

	if redirect_uri != "" {
		redirect = redirect_uri
	} else {
		redirect = api.redirect_uri
	}
	if redirect == "" { // Check Redirect
		panic(&APIError{when: time.Now(), error_code: "21305", message: "Parameter absent: redirect_uri"})
	}
	var params = map[string]interface{}{
		"client_id":     api.app_key,
		"client_secret": api.app_secret,
		"redirect_uri":  redirect,
		"code":          code,
		"grant_type":    "authorization_code",
	}
	result = httpCall(the_url, HTTP_POST, "", params)

	return result
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
		log.Fatal(err)
		os.Exit(1)
	}
}
