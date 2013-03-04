package weibo

import (
	// "bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
)

// func TestHttpCall(t *testing.T) {
// 	api := NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php")
// 	api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1519925461)
// 	// kws := map[string]string{
// 	// 	"uid": "2684726573",
// 	// }
// 	update_status := map[string]string{
// 		"status": "Test Go SDK",
// 	}
// 	api.Post.Call("statuses/update", update_status)
// 	// api.Get.Call("statuses/user_timeline", kws)
// }

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
		"status": "Test Go SDK",
		"pic":    pic,
	}
	fmt.Println(reflect.TypeOf(pic))
	api.Upload.Call("statuses/upload", update_status)

}
