# select

## 管道的读写
#### select 只能作用于管道，包括输入和输出
#### select 每个case只能操作一个管道。如果管道没有数据读取，则阻塞；如果没有缓冲区写入，会阻塞；若没有符合条件的case，则会一直阻塞

## 返回值
#### select case语句最多给两个变量赋值
```go
package main
import "fmt"
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
	}
}


```
#### case读操作触发条件：读到数据或者管道没有数据且已经关闭

## default
#### default不能处理管道读写，若所有case都阻塞，会走default语句


## 使用案例
#### 永久阻塞：启动协程处理任务时，不希望main函数退出，则main函数需要永久阻塞
#### 快速检错：管道用来传输错误信息，若case读取不到错误内容，说明没有错误
#### 限时等待：select可以创建只有一定时效的管道
```go
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
	case <- time.After(5 * time.Second):
		fmt.Println("it is time to end")
	}
}
```

## 实现原理 
func selectgo(cas0 *scase, order0 *unit16, ncase int) (int, bool)
#### 若多个case满足条件，则随机选择一个。原因是：源码中selectgo()函数会把原始case顺序打乱，选择的时候是随机的，而且orders数组后半部分会存在管道的加锁顺序，避免重复加锁
#### case语句中，若管道值为nil,向管道写数据的时候，不会出发panic，而是阻塞。原因：源码中会对管道为nil的case语句进行标记，永远不会命中，所以不会panic
#### 返回值为第一个表示case在代码中出现的编号，第二个表示是否从管道中读取到了数据。

## 小结
## 使用select读取管道时，尽量检查读取是否成功，方便发现管道异常
## select只能操作管道
## 每个case只能操作一个管道，要么读，要么写
## 多个case的执行顺序是随机
## 存在default语句时，select不会阻塞