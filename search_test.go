/**
 * Author: Tony.Shao(xiocode@gmail.com)
 * Date: 13-03-15
 * Version: 0.02
 */
package weigo

import (
	// "fmt"
	"testing"
)

func Test_GET_search_topics(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"q": "肖申克的救赎",
	}
	result := new(Topic)
	err := api.GET_search_topics(kws, result)
	debugCheckError(err)
	debugPrintln(*result)
}
