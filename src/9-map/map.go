package main

import (
	"fmt"
)

func main() {
	// 声明
	var myMap map[string]string

	if myMap == nil {
		fmt.Println("nil map")
	}

	myMap = make(map[string]string, 10)

	myMap["one"] = "java"
	myMap["two"] = "java"
	myMap["three"] = "java"

	fmt.Println(myMap)

	myMap2 := make(map[string]string, 2)
	myMap2["a"] = "java"
	myMap2["b"] = "java"
	myMap2["c"] = "java"
	fmt.Println(myMap2)

	myMap3 := map[string]string{
		"a": "java",
		"b": "java",
	}

	fmt.Println(myMap3)

	cityMap := make(map[string]string)

	cityMap["china"] = "beijing"
	cityMap["usa"] = "NewYork"

	for key, value := range cityMap {
		fmt.Println(key, " ", value)
	}

	delete(cityMap, "china")
	fmt.Println(cityMap)

	cityMap["china"] = "tianjin"
	fmt.Println(cityMap)

}
