package main

/*
下面代码输出什么？
答案：default
解析：ch为nil，读写会阻塞，所以返回default
*/
func main() {
	var ch chan int
	select {
	case v, ok := <-ch:
		println(v, ok)
	default:
		println("default")
	}
}
