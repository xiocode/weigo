Go Weibo SDK
========
Sina Weibo SDK For Gopher

文档请看测试用例，哈哈！
开玩笑，剩下的回头补上来,有问题欢迎提！

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
)

func main() {
	api := weigo.NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php")
	api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1519925461)

	//Update Status
	update_status := map[string]interface{}{
		"status": "Test Go Weibo SDK 2",
	}
	result1 := api.Post.Call("statuses/update", update_status)
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
	result2 := api.Upload.Call("statuses/upload", update_status_pic)
	fmt.Println(result2)

	kws := map[string]interface{}{
		"uid": "2684726573",
	}
	result3 := api.Get.Call("statuses/user_timeline", kws)
	fmt.Println(result3)

}

```

Weibo: http://weibo.com/xceman
Gmail: xiocode@gmail.com