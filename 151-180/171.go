package main

/*
下面代码有什么问题，请说明？
答案：打印0123456789后，会死循环
*/
import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	for {
	}
}
