package main

import (
	"fmt"
	"time"
)

func SelectAssign(c chan string) {

	select {
	case <-c:
		fmt.Println("0")
	case d := <-c:
		fmt.Println(d)
	case d, ok := <-c:
		if !ok {
			fmt.Println("no data found")
		} else {
			fmt.Println("receive :", d)
		}
	case <-time.After(5 * time.Second):
		fmt.Println("it is time to end")
	}
}

func main() {

	var c = make(chan string, 10)
	//c <- "1"
	close(c)

	SelectAssign(c)

}
