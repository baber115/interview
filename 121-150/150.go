package main

import "fmt"

/*
下面的代码输出什么？

答案：7 6 8
*/
func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func main() {
	f := F(5)
	defer func() { // A
		fmt.Println(f(), "c")
	}()
	defer fmt.Println(f(), "b") // B
	i := f()
	fmt.Println(i, "a")
}
