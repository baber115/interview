package main

import "fmt"

// 关于channel，下面正确的是
/**
A. var ch chan int
B. ch := make(chan int)
C. <-ch
D. ch<-

答案
ABC
AB，是定义一个channel
CD，分别是取值和存值，但是存值必须有值，所以D是错误的
*/

func main() {
	ch := make(chan int, 1)
	ch <- 1
	a := <-ch
	close(ch)
	fmt.Println(a)
}
