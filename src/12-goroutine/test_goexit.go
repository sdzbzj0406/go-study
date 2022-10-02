package main

import (
	"fmt"
	"time"
)

func main2() {
	go func() {

		defer fmt.Println("a.defer")

		//return

		func() {
			defer fmt.Println("b.defer")

			return

			// 退出整个go进程，不会执行a
			// runtime.Goexit()
			fmt.Println("b")
		}()

		fmt.Println("a")

		return
	}()

	for true {
		time.Sleep(1 * time.Second)
	}

}
