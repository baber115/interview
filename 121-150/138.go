package main

import "fmt"

/*
下面的代码输出什么？
答案：3 2 1
解析：
第一次循环，写操作准备好，执行o(3)，输出3；
第二次循环，读操作准备好，执行o(2)，输出2，并把c赋值为nil
第三次循环，c=nil，default输出1
*/
var o = fmt.Print

func main() {
	c := make(chan int, 1)
	for range [3]struct{}{} {
		select {
		default:
			o(1)
		case <-c:
			o(2)
			c = nil
		case c <- 1:
			o(3)
		}
	}
}
