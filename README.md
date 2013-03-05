Go Weibo SDK
========
Sina Weibo SDK For Gopher

文档请看测试用例，哈哈！
开玩笑，剩下的回头补上来,有问题欢迎提！

##Install:
```go
go get -u github.com/xiocode/weigo
```

##Useage:
http://open.weibo.com/wiki/API%E6%96%87%E6%A1%A3_V2
参照官方文档调用对应的方法
```go
package main

import (
	// "bytes"
	"fmt"
	"github.com/xiocode/weigo"
	"os"
	// "time"
)

func main() {
	api := weigo.NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php", "code")

	///////////////////////////////////////////////授权/////////////////////////////////////////////////////////
	// authorize_url := api.GetAuthorizeUrl("", map[string]interface{}{"force_login": 1})
	authorize_url, err1 := api.GetAuthorizeUrl(nil)
	if err != nil {
		fmt.Println(err1)
	}
	fmt.Println(authorize_url) // 授权地址
	//浏览器访问授权地址，获取CODE
	result, err2 := api.RequestAccessToken("15fc75c174ccc6e81f2f5060d0555d48") //使用返回的CODE请求授权
	if err != nil {
		fmt.Println(err2)
	}
	fmt.Println(result) // 授权信息
	access_token := result["result"].(string)
	expires := result["result"].(float64)

	///////////////////////////////////////////////方法调用/////////////////////////////////////////////////////////

	api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1519925461)
	//Update Status
	update_status := map[string]interface{}{
		"status": "Test Go Weibo SDK 3",
	}
	result1, err1 := api.Post.Call("statuses/update", update_status)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(result1)

	//Upload Pic & Update Status
	pic, err := os.Open("test.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	update_status_pic := map[string]interface{}{
		"status": "Test Go Weibo SDK Upload Picture & Update Status 2",
		"pic":    pic,
	}
	result2, err2 := api.Upload.Call("statuses/upload", update_status_pic)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(result2)

	kws := map[string]interface{}{
		"uid": "2684726573",
	}
	result3, err3 := api.Get.Call("statuses/user_timeline", kws)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(result3)

}
```

Weibo: http://weibo.com/xceman @XIOCODE
Gmail: xiocode@gmail.com