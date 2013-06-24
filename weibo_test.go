package weigo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetAuthorizeUrl(t *testing.T) {
	t.SkipNow()
	authorize_url, err := api.GetAuthorizeUrl(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(authorize_url)
}

func TestRequestAccessToken(t *testing.T) {
	t.SkipNow()
	var result map[string]interface{}
	err := api.RequestAccessToken("1fdaa295b73d2a9568e284383ced5e9e", &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	access_token := result["access_token"]
	fmt.Println(reflect.TypeOf(access_token), access_token)
	expires_in := result["expires_in"]
	fmt.Println(reflect.TypeOf(expires_in), expires_in)
}
