package main

/*
下面代码输出什么？

答案：会报错
解析：A协程起的时候主协程已经把channel关闭了，A在往里放值会报错
*/
import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 100)
	// A
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	// B
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 10)
}
