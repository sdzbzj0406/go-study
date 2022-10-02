package main

import (
	"fmt"
	"time"
)

func main2() {

	c := make(chan int, 3)

	fmt.Println(len(c), " ", cap(c))

	go func() {
		defer fmt.Println("go defer")

		// 写满会阻塞
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("go running")
			fmt.Println(i, " ", len(c), " ", cap(c))
		}
		fmt.Println("go end")

	}()

	time.Sleep(2 * time.Second)

	// 读多会阻塞
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num: ", num)
	}

	fmt.Println("main end")
}
