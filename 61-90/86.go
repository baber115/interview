package main

import (
	"fmt"
	"runtime"
)

/*
下面代码会触发异常吗？请说明。

答案：可能会也可能不会

1、有多个channel发生操作时，select会随机选择其中一个执行
2、select需要default，如果没有default的话，select会一直阻塞
3、select的每一个case，必须是channel操作
*/
func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
