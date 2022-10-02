package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0

	for {
		i++
		fmt.Println("new task  i=", i)
		time.Sleep(1 * time.Second)
	}

}

func main1() {

	go newTask()

	// main进程关闭后，go进程也会同步停掉
	time.Sleep(3 * time.Second)

	return

	//i := 0
	//for {
	//	i++
	//	fmt.Println("main func  i=", i)
	//	time.Sleep(1 * time.Second)
	//}

}
