package weibo

import (
	"testing"
)

func TestHttpCall(t *testing.T) {
	api := NewAPIClient("", "", "")
	api.SetAccessToken("Token", 99999999999999)
	kws := map[string]string{
		"uid": "1642634100",
	}

	api.Get.Call("statuses/user_timeline", kws)
}
