package weigo

import (
	"testing"
)

func Test_GET_favorites(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"access_token": api.access_token,
		"source":       api.app_key,
	}
	result := new(Favorites)
	err := api.GET_favorites(kws, result)
	debugCheckError(err)
	debugPrintln(*((*result.Favorites)[0]).Status)
}
