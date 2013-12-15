package weigo

import (
	"testing"
)

func Test_GET_suggestions_statuses_reorder(t *testing.T) {
	// t.SkipNow()
	kws := map[string]interface{}{
		"section": 60,
	}
	result := new(Topic)
	err := api.GET_suggestions_statuses_reorder(kws, result)
	debugCheckError(err)
	debugPrintln(*result)
}
