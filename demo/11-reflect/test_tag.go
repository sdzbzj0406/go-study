package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"name`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {

	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagString := t.Field(i).Tag.Get("info")
		fmt.Println(tagString)
	}

}

func main2() {

	var name resume

	findTag(&name)

}
