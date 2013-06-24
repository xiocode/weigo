package weigo

import (
	"testing"
)

func Test_GET_comments_show(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"id": "3551749023600582",
	}
	result := new(Comments)
	err := api.GET_comments_show(kws, result)
	debugCheckError(err)
	debugPrintln(len(*result.Comments))
}

func Test_POST_comments_create(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"id":      "3551749023600582",
		"comment": "Testing...Testing...",
	}
	result := new(Comment)
	err := api.POST_comments_create(kws, result)
	debugCheckError(err)
	debugPrintln(*result)
}
