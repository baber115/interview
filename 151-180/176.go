package main

import (
	"fmt"
	"sync"
)

/*
下面的代码有什么问题？
答案：wg.Add后没有wg.Done，会panic
解析：WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量。
Add(n) 把计数器设置为n ，Done() 每次把计数器-1 ，wait() 会阻塞代码的运行，直到计数器地值减为0。
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("1")
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}
