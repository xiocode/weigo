package weigo

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestHttpCallPost(t *testing.T) {
	api := NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php")
	api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1519925461)
	// kws := map[string]string{
	// 	"uid": "2684726573",
	// }
	update_status := map[string]interface{}{
		"status": "Test Go Weibo SDK",
	}
	api.Post.Call("statuses/update", update_status)
	// api.Get.Call("statuses/user_timeline", kws)
}

func TestHttpCallUpload(t *testing.T) {

	api := NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php")
	api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1519925461)
	// kws := map[string]string{
	// 	"uid": "2684726573",
	// }
	pic, err := os.Open("test.jpg")
	if err != nil {
		fmt.Println(err)
	}
	update_status := map[string]interface{}{
		"status": "Test Go Weibo SDK Upload Picture & Update Status",
		"pic":    pic,
	}
	fmt.Println(reflect.TypeOf(pic))
	api.Upload.Call("statuses/upload", update_status)

}

func TestHttpCallGet(t *testing.T) {
	api := NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php")
	api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1519925461)
	authorize_url := api.GetAuthorizeUrl("", map[string]interface{}{"force_login": 1})
	kws := map[string]interface{}{
		"uid": "2684726573",
	}
	api.Get.Call("statuses/user_timeline", kws)
}

func TestHttpCallRequestToken(t *testing.T) {
	api := NewAPIClient("3417104247", "f318153f6a80329f06c1d20842ee6e91", "http://127.0.0.1/callback")
	// authorize_url := api.GetAuthorizeUrl("", map[string]interface{}{"force_login": 1})
	authorize_url := api.GetAuthorizeUrl("", nil)
	fmt.Println(authorize_url)
	token := api.RequestAccessToken("15fc75c174ccc6e81f2f5060d0555d48", "http://127.0.0.1/callback")
	fmt.Println(token)
}
