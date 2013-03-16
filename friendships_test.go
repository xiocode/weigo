/**
 * Author: Tony.Shao(xiocode@gmail.com)
 * Date: 13-03-15
 * Version: 0.02
 */
package weigo

import (
	"testing"
)

var api *APIClient

func init() {
	if api == nil {
		api = NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php", "code")
		api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1519925461)
	}
}

func Test_GET_users_show(t *testing.T) {
	kws := map[string]interface{}{
		"uid": "2684726573",
	}
	result := new(Friendships)
	err := api.GET_friendships_friends(kws, result)
	debugCheckError(err)
	debugPrintln(len(*result.Users))
}

func Test_GET_users_counts(t *testing.T) {
	kws := map[string]interface{}{
		"uid":  "1580095602",
		"suid": "2684726573",
	}
	result := new(Friendships)
	err := api.GET_friendships_friends_in_common(kws, result)
	debugCheckError(err)
	debugPrintln(len(*result.Users))
}
