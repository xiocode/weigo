package weigo

import (
	// "bytes"
	"fmt"
	// "os"
	"reflect"
	"testing"
)

// func TestHttpCallPost(t *testing.T) {
// 	api := NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php", "code")
// 	api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1519925461)
// 	// kws := map[string]string{
// 	// 	"uid": "2684726573",
// 	// }
// 	update_status := map[string]interface{}{
// 		"status": "Test Go Weibo SDK",
// 	}
// 	result, err := api.Post.Call("statuses/update", update_status)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// 	// api.Get.Call("statuses/user_timeline", kws)
// }

// func TestHttpCallUpload(t *testing.T) {

// 	api := NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php", "code")
// 	api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1519925461)
// 	// kws := map[string]string{
// 	// 	"uid": "2684726573",
// 	// }
// 	pic, err := os.Open("test.jpg")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	update_status := map[string]interface{}{
// 		"status": "Test Go Weibo SDK Upload Picture & Update Status with http_length",
// 		"pic":    pic,
// 	}
// 	result, err := api.Upload.Call("statuses/upload", update_status)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// }

// func TestHttpCallGet(t *testing.T) {
// 	api := NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php", "code")
// 	api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1519925461)
// 	// authorize_url, err := api.GetAuthorizeUrl(map[string]interface{}{"force_login": 1})
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	kws := map[string]interface{}{
// 		"uid": "2684726573",
// 	}
// 	// var result map[string]interface{}
// 	result := new(Timeline)
// 	err := api.GET_statuses_user_timeline(kws, &result)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// 	fmt.Println(*result.Statuses)
// 	fmt.Println(len(*result.Statuses))
// }

func TestGetAuthorizeUrl(t *testing.T) {
	api := NewAPIClient("3417104247", "f318153f6a80329f06c1d20842ee6e91", "http://127.0.0.1/callback", "code")
	authorize_url, err := api.GetAuthorizeUrl(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(authorize_url)
}

func TestRequestAccessToken(t *testing.T) {
	api := NewAPIClient("3417104247", "f318153f6a80329f06c1d20842ee6e91", "http://127.0.0.1/callback", "code")
	var result map[string]interface{}
	err := api.RequestAccessToken("1fdaa295b73d2a9568e284383ced5e9e", &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	access_token := result["access_token"]
	fmt.Println(reflect.TypeOf(access_token), access_token)
	expires_in := result["expires_in"]
	fmt.Println(reflect.TypeOf(expires_in), expires_in)
}
