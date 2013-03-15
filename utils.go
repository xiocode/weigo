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

func JSONParser(body string, result interface{}) (err error) {
	body_bytes := []byte(body)
	err = json.Unmarshal(body_bytes, result)
	if err != nil {
		return
	}
	return nil
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
