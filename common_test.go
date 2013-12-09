package weigo

import (
	"testing"
)

func Test_GET_common_get_province(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"country":      "001",
		"access_token": api.access_token,
		"source":       api.app_key,
	}
	result := new([]map[string]string)
	err := api.GET_common_get_province(kws, result)
	debugCheckError(err)
	debugPrintln(((*result)[0]))
}
