package weigo

import (
	"testing"
	"time"
)

var api *APIClient

type Times struct {
	time.Time
}

func init() {
	if api == nil {
		api = NewAPIClient("3231340587", "702b4bcc6d56961f569943ecee1a76f4", "http://2.xweiboproxy.sinaapp.com/callback.php", "code")
		api.SetAccessToken("2.00VBqgvCZS4gWDb3940dd56eFfitSB", 1524137445)
	}
}

func Test_GET_statuses_user_timeline(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"mid": "yybyOssa9",
		// "type":     1,
		// "isBase62": 1,
	}
	var result string
	err := api.GET_statuses_queryid(kws, &result)
	debugCheckError(err)
	debugPrintln(result)
}

// func Test_GET_statuses_home_timeline(t *testing.T) {
// 	kws := map[string]interface{}{
// 		"uid": "2684726573",
// 	}
// 	result := new(Statuses)
// 	err := api.GET_statuses_home_timeline(kws, result)
// 	debugCheckError(err)
// 	debugPrintln(len(*result.Statuses))
// }

// func Test_GET_statuses_repost_timeline(t *testing.T) {
// 	kws := map[string]interface{}{
// 		"id": "3551749023600582",
// 	}
// 	result := new(Reposts)
// 	err := api.GET_statuses_repost_timeline(kws, result)
// 	debugCheckError(err)
// 	debugPrintln(len(*result.Reposts))
// }

// func Test_POST_statuses_repost(t *testing.T) {
// 	kws := map[string]interface{}{
// 		"id": "3551749023600582",
// 	}
// 	result := new(Status)
// 	err := api.POST_statuses_repost(kws, result)
// 	debugCheckError(err)
// 	debugPrintln(*result)
// }

// func Test_POST_statuses_repost(t *testing.T) {
// 	kws := map[string]interface{}{
// 		"status": "Testing...Testing...",
// 	}
// 	result := new(Status)
// 	err := api.POST_statuses_update(kws, result)
// 	debugCheckError(err)
// 	debugPrintln(*result)
// }

// func Test_POST_statuses_repost(t *testing.T) {
// 	kws := map[string]interface{}{
// 		"id": "3556138715301190",
// 	}
// 	result := new(Status)
// 	err := api.POST_statuses_destroy(kws, result)
// 	debugCheckError(err)
// 	debugPrintln(*result)
// }
