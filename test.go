package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

var str = `{
	"a":"1",
	"data": {
		"user": "gogo"
	}
}`

func main() {
	var x struct{ User string  `json:"user"` }
	result := gjson.Get(str, "data")
	err := json.Unmarshal([]byte(result.Raw), &x)
	if nil != err {
		panic(err)
	}
	fmt.Println(x.User)
}
