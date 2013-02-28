package weibo

import (
	"testing"
)

func TestHttpCall(t *testing.T) {
	api := NewAPIClient("1", "2", "3")
	kws := map[string]string{
		"a":     "1",
		"b":     "a",
		"count": "123",
	}
	api.Post.Call("statuses/public_timeline", kws)
}
