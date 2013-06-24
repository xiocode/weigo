/**
 * Author: Tony.Shao(xiocode@gmail.com)
 * Date: 13-03-15
 * Version: 0.02
 */
package weigo

import (
	"testing"
)

func Test_GET_friendships_friends(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"uid": "2684726573",
	}
	result := new(Friendships)
	err := api.GET_friendships_friends(kws, result)
	debugCheckError(err)
	debugPrintln(len(*result.Users))
}

func Test_GET_friendships_friends_in_common(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"uid":  "1580095602",
		"suid": "2684726573",
	}
	result := new(Friendships)
	err := api.GET_friendships_friends_in_common(kws, result)
	debugCheckError(err)
	debugPrintln(len(*result.Users))
}
