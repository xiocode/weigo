package weigo

import (
	// "bytes"
	"fmt"
	"os"
	// "reflect"
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

func TestHttpCallUpload(t *testing.T) {

	api := NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php", "code")
	api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1519925461)
	// kws := map[string]string{
	// 	"uid": "2684726573",
	// }
	pic, err := os.Open("test.jpg")
	if err != nil {
		fmt.Println(err)
	}
	update_status := map[string]interface{}{
		"status": "Test Go Weibo SDK Upload Picture & Update Status with http_length",
		"pic":    pic,
	}
	result, err := api.Upload.Call("statuses/upload", update_status)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

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
// 	result, err := api.Get.Call("statuses/user_timeline", kws)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// }

// func TestHttpCallRequestToken(t *testing.T) {
// 	api := NewAPIClient("3417104247", "f318153f6a80329f06c1d20842ee6e91", "http://127.0.0.1/callback", "code")
// 	// authorize_url := api.GetAuthorizeUrl("", map[string]interface{}{"force_login": 1})
// 	// authorize_url, err := api.GetAuthorizeUrl(nil)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// fmt.Println(authorize_url)

// 	result, err := api.RequestAccessToken("d6757d781936933dd6184b6b4e5143aa")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// 	access_token := result["access_token"]
// 	fmt.Println(reflect.TypeOf(access_token), access_token)
// 	expires_in := result["expires_in"]
// 	fmt.Println(reflect.TypeOf(expires_in), expires_in)
// }
