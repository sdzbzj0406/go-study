package main

import (
	"fmt"
	"reflect"
)

func main1() {
	var a string

	// pair <string, abcd>
	a = "hello"

	// pair<type:, value>
	var allType interface{}
	allType = a

	value, _ := allType.(string)
	fmt.Println(value)

	fmt.Println(reflect.ValueOf(allType))
	fmt.Println(reflect.TypeOf(allType))

}
