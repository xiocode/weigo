/**
 * Author: Tony.Shao(xiocode@gmail.com)
 * Date: 13-03-15
 * Version: 0.02
 */
package weigo

import (
	"testing"
)

func Test_GET_users_show(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"uid": "2684726573",
	}
	result := new(User)
	err := api.GET_users_show(kws, result)
	debugCheckError(err)
	debugPrintln(*result)
}

func Test_GET_users_counts(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"uids": "1580095602,2684726573",
	}
	result := new([]UserCounts)
	err := api.GET_users_counts(kws, result)
	debugCheckError(err)
	debugPrintln(*result)
}
