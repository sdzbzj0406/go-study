package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string `json:"title"`
	Year   int
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {

	m := Movie{"喜剧之王", 2000, 10, []string{"aaa", "bbb"}}

	// 对象转json
	j, error := json.Marshal(m)

	if error != nil {
		fmt.Println("error")
	} else {
		fmt.Printf("%s", j)
	}

	fmt.Println()
	// json转回对象
	myMovie := Movie{}
	err := json.Unmarshal(j, &myMovie)
	if err != nil {
		fmt.Println("has error")
	} else {
		fmt.Printf("%v", myMovie)
	}

}
