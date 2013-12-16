/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          weibo.go
 * Description:   weigo core
 */
package weigo

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/going/toolkit/simplejson"
	"github.com/going/toolkit/to"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	HTTP_GET    int = 0
	HTTP_POST   int = 1
	HTTP_UPLOAD int = 2
)

func httpCall(the_url string, method int, authorization string, params map[string]interface{}) (body string, err error) {
	var url_params string
	var multipart_data *bytes.Buffer //For Upload Image
	var http_url string
	var http_body io.Reader
	var content_type string
	var request *http.Request
	var HTTP_METHOD string

	switch method {
	case HTTP_GET:
		HTTP_METHOD = "GET"
		url_params, err = encodeParams(params)
		http_url = fmt.Sprintf("%v?%v", the_url, url_params)
		http_body = nil
	case HTTP_POST:
		HTTP_METHOD = "POST"
		url_params, err = encodeParams(params)
		content_type = "application/x-www-form-urlencoded"
		http_url = the_url
		http_body = strings.NewReader(url_params)
	case HTTP_UPLOAD:
		HTTP_METHOD = "POST"
		the_url = strings.Replace(the_url, "https://api.", "https://upload.api.", 1)
		content_type, multipart_data, err = encodeMultipart(params)
		http_url = the_url
		http_body = multipart_data
	}
	if err != nil {
		return "", err
	}

	client := new(http.Client)
	request, err = http.NewRequest(HTTP_METHOD, http_url, http_body)
	if err != nil {
		return "", err
	}
	request.Header.Add("Accept-Encoding", "gzip")

	switch method {
	case HTTP_POST:
		request.Header.Add("Content-Type", content_type)
	case HTTP_UPLOAD:
		request.Header.Add("Content-Type", content_type)
		request.Header.Add("Content-Length", to.String(multipart_data.Len()))
	}
	if authorization != "" {
		request.Header.Add("Authorization", fmt.Sprintf("OAuth2 %s", authorization))
	}

	response, err := client.Do(request) // Do Request
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err = read_body(response)
	if err != nil {
		return "", err
	}

	return body, nil
}

func encodeParams(params map[string]interface{}) (string, error) {
	if len(params) > 0 {
		values := url.Values{}
		for key, value := range params {
			values.Add(key, to.String(value))
		}
		return values.Encode(), nil
	}
	return "", errors.New("Params Is Empty!")
}

func encodeMultipart(params map[string]interface{}) (multipartContentType string, multipartData *bytes.Buffer, err error) {
	if len(params) > 0 {
		multipartData := new(bytes.Buffer)
		bufferWriter := multipart.NewWriter(multipartData) // type *bytes.Buffer
		defer bufferWriter.Close()
		var multipartContentType string
		for key, value := range params {
			switch value.(type) {
			case *os.File:
				picdata, err := bufferWriter.CreateFormFile(key, value.(*os.File).Name())
				if err != nil {
					return "", nil, err
				}
				multipartContentType = bufferWriter.FormDataContentType()
				io.Copy(picdata, value.(*os.File))
			default:
				bufferWriter.WriteField(key, to.String(value))
			}
		}
		return multipartContentType, multipartData, nil
	}
	return "", nil, errors.New("Params Is Empty!")
}

func read_body(response *http.Response) (string, error) {

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(response.Body)
		if err != nil {
			return "", err
		}
		defer reader.Close()
		contents, err := ioutil.ReadAll(reader)
		if err != nil {
			return "", err
		}
		return string(contents), nil
	default:
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		return string(contents), nil
	}

	return "", errors.New("Unknow Errors")
}

type HttpObject struct {
	client *APIClient
	method int
}

func (http *HttpObject) call(uri string, params map[string]interface{}, result interface{}) error {
	if http.client.is_expires() {
		return &APIError{When: time.Now(), ErrorCode: 21327, Message: "expired_token"}
	}

	url := fmt.Sprintf("%s%s.json", http.client.api_url, uri)
	body, err := httpCall(url, http.method, http.client.access_token, params)
	if err != nil {
		return err
	}
	fmt.Println(body)
	if strings.Trim(body, " ") == "" {
		return errors.New("Nothing Return From Http Requests!")
	}
	jsonbody, err := simplejson.NewJson([]byte(body))
	if err != nil {
		return err
	}
	_, ok := jsonbody.CheckGet("error_code")
	if ok {
		errcode, _ := jsonbody.Get("error_code").Int64()
		errmessage, _ := jsonbody.Get("error").String()
		err := &APIError{When: time.Now(), ErrorCode: errcode, Message: errmessage}
		return err
	}

	if json.Unmarshal([]byte(body), result); err != nil {
		return err
	}
	return nil
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
	get           *HttpObject
	post          *HttpObject
	upload        *HttpObject
}

func (api *APIClient) is_expires() bool {
	return api.access_token == "" || api.expires < time.Now().Unix()
}

func NewAPIClient(app_key, app_secret, redirect_uri, response_type string) *APIClient {
	api := &APIClient{
		app_key:       app_key,
		app_secret:    app_secret,
		redirect_uri:  redirect_uri,
		response_type: response_type,
		domain:        "api.weibo.com",
		version:       "2",
	}

	api.auth_url = fmt.Sprintf("https://%s/oauth2/", api.domain)
	api.api_url = fmt.Sprintf("https://%s/%s/", api.domain, api.version)
	api.get = &HttpObject{client: api, method: HTTP_GET}
	api.post = &HttpObject{client: api, method: HTTP_POST}
	api.upload = &HttpObject{client: api, method: HTTP_UPLOAD}
	return api
}

func (api *APIClient) SetAccessToken(access_token string, expires int64) *APIClient {
	api.access_token = access_token
	api.expires = expires
	return api
}

func (api *APIClient) GetAuthorizeUrl(params map[string]interface{}) (string, error) {

	url_params := map[string]interface{}{
		"client_id":     api.app_key,
		"response_type": api.response_type,
		"redirect_uri":  api.redirect_uri,
	}
	for key, value := range params {
		url_params[key] = value
	}
	encode_params, err := encodeParams(url_params)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s?%s", api.auth_url, "authorize", encode_params), nil
}

func (api *APIClient) RequestAccessToken(code string, result *map[string]interface{}) error {
	the_url := fmt.Sprintf("%s%s", api.auth_url, "access_token")
	params := map[string]interface{}{
		"client_id":     api.app_key,
		"client_secret": api.app_secret,
		"redirect_uri":  api.redirect_uri,
		"code":          code,
		"grant_type":    "authorization_code",
	}
	body, err := httpCall(the_url, HTTP_POST, "", params)
	if err != nil {
		return err
	}
	if json.Unmarshal([]byte(body), result); err != nil {
		return err
	}

	return nil
}

func (a *APIClient) GET(uri string, params map[string]interface{}, result interface{}) error {
	return a.get.call(uri, params, result)
}

func (a *APIClient) POST(uri string, params map[string]interface{}, result interface{}) error {
	return a.post.call(uri, params, result)
}

func (a *APIClient) UPLOAD(uri string, params map[string]interface{}, result interface{}) error {
	return a.upload.call(uri, params, result)
}

type APIError struct {
	When      time.Time
	ErrorCode int64
	Message   string
}

func (err *APIError) Error() string {
	if err == nil {
		return "Error with unknown reason"
	}
	return fmt.Sprintf("APIError When: %v ErrorMessage: %v ErrorCode: %v", err.When, err.Message, err.ErrorCode)
}
