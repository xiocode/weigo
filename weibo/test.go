package weibo

import (
	"fmt"
	"net/http"
	"reflect"
)

func a() *http.Client {
	s := new(http.Client)
	fmt.Println(reflect.TypeOf(s))
	return s
}

func main() {
	var client *http.Client
	client = a()
	fmt.Println(reflect.TypeOf(client))
}
