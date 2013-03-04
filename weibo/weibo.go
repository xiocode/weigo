/**
 * Author: xio
 * Date: 13-2-27
 * Version: 0.01
 */
package weibo

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
	"time"
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

func httpCall(the_url string, method int, authorization string, kws map[string]interface{}) bool {
	var params string
	var multipart *bytes.Buffer
	var http_url string
	var http_body string
	var contentType string
	var multipartContentType string
	var request *http.Request
	var HTTP_METHOD string
	var err error

	if method == HTTP_UPLOAD {
		the_url = strings.Replace(the_url, "https://api.", "https://upload.api.", 1)
		multipartContentType, multipart = encodeMultipart(kws)
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

	switch method {
	case HTTP_GET:
		HTTP_METHOD = "GET"
		contentType = ""
	case HTTP_POST:
		HTTP_METHOD = "POST"
		contentType = "application/x-www-form-urlencoded"
	case HTTP_UPLOAD:
		HTTP_METHOD = "POST"
		contentType = multipartContentType
	}

	client := new(http.Client)
	if method == HTTP_UPLOAD {
		request, err = http.NewRequest(HTTP_METHOD, http_url, multipart) //Make New Request
	} else {
		request, err = http.NewRequest(HTTP_METHOD, http_url, strings.NewReader(http_body)) //Make New Request
	}
	request.Header.Add("Accept-Encoding", "gzip")
	request.Header.Add("Content-Type", contentType)

	if authorization != "" {
		request.Header.Add("Authorization", fmt.Sprintf("OAuth2 %s", authorization))
	}

	response, err := client.Do(request) // Do Request
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	checkError(err)
	defer response.Body.Close()

	body := read_body(response)
	parse_json(body)
	return true
}

/************************************************************
*
* Encode Params *
*
*************************************************************/
func encodeParams(kws map[string]interface{}) (params string) {
	if len(kws) > 0 {
		values := url.Values{}
		for key, value := range kws {
			switch value.(type) {
			case string:
				values.Add(key, value.(string))
			case int:
				values.Add(key, string(value.(int)))
			}
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
func encodeMultipart(kws map[string]interface{}) (multipartContentType string, multipartData *bytes.Buffer) {
	if len(kws) > 0 {
		multipartData = new(bytes.Buffer)
		bufferWriter := multipart.NewWriter(multipartData)
		for key, value := range kws {
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
				fmt.Println("DEBUG")
				fmt.Println(picdata)
			}
		}
	}
	return multipartContentType, multipartData
}

/************************************************************
*
* Read Response Body *
*
*************************************************************/
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

/************************************************************
*
*Parse Json To JSON Struct *
*
*************************************************************/
func parse_json(body string) (result map[string]interface{}) {
	// jsonDict := new(map[string]interface{})
	// data, err := json.NewDecoder(strings.NewReader(body))
	// checkError(err)
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
func (http *HttpObject) Call(uri string, kws map[string]interface{}) {
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

/************************************************************
*
* Set User AccessToken & Expires *
*
*************************************************************/

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
		log.Fatal(err)
		os.Exit(1)
	}
}
