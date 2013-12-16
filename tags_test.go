package weigo

import (
	"testing"
)

func Test_GET_tags(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"uid": "2684726573",
	}
	result := new([]map[string]interface{})
	err := api.GET_tags(kws, result)
	debugCheckError(err)
	debugPrintln(((*result)[0]["weight"]))
}

func Test_GET_tags_tags_batch(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"uids": "2684726573",
	}
	result := new([]Tags)
	err := api.GET_tags_tags_batch(kws, result)
	debugCheckError(err)
	debugPrintln((*result))
}
