/**
 * Author: Tony.Shao(xiocode@gmail.com)
 * Date: 13-02-27
 * Version: 0.02
 */

package weigo

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func decode(body string, instance interface{}) (err error) {
	b := []byte(body)
	err = json.Unmarshal(b, instance)
	if err != nil {
		return
	}
	return nil
}

func checkError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func debugPrintln(message ...interface{}) {
	fmt.Println(message)
}

func debugTypeof(element interface{}) interface{} {
	return reflect.TypeOf(element)
}

func debugCheckError(err error) {
	if err != nil {
		debugPrintln(err)
	}
}
