package weigo

import (
	"testing"
)

func Test_GET_account_privacy(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"access_token": api.access_token,
		"source":       api.app_key,
	}
	result := new(Config)
	err := api.GET_account_privacy(kws, result)
	debugCheckError(err)
	debugPrintln((*result))
}

func Test_GET_account_rate_limit_status(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"access_token": api.access_token,
		"source":       api.app_key,
	}
	result := new(LimitStatus)
	err := api.GET_account_rate_limit_status(kws, result)
	debugCheckError(err)
	debugPrintln((*result))
}

func Test_GET_account_get_uid(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"access_token": api.access_token,
		"source":       api.app_key,
	}
	result := new(UserID)
	err := api.GET_account_get_uid(kws, result)
	debugCheckError(err)
	debugPrintln((*result))
}

func Test_GET_account_profile_school_list(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"access_token": api.access_token,
		"source":       api.app_key,
		"keyword":      "科大",
	}
	result := new([]School)
	err := api.GET_account_profile_school_list(kws, result)
	debugCheckError(err)
	debugPrintln((*result))
}

func Test_GET_account_get_email(t *testing.T) {
	t.SkipNow()
	kws := map[string]interface{}{
		"access_token": api.access_token,
		"source":       api.app_key,
	}
	result := new(Email)
	err := api.GET_account_get_email(kws, result)
	debugCheckError(err)
	debugPrintln((*result))
}
