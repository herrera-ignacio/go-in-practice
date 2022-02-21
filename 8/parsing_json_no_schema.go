package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var myJson = []byte(`{
	"firstName": "John",
	"lastName": "Smith",
	"age": 25,
	"children": [
		"Mary",
		"Moe",
		"Jack"
	]
}`)

func printJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)
	case float64:
		fmt.Println("is float64", vv)
	case []interface{}:
		fmt.Println("is an array", vv)
		for i, u := range vv {
			fmt.Println(i, " ")
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object", vv)
		for i, u := range vv {
			fmt.Println(i, " ")
			printJSON(u)
		}
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	var f interface{}
	err := json.Unmarshal(myjson, &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(f)
}
