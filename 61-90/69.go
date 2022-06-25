package main

import "fmt"

/**
以下代码输出什么

答案：7
一、先运行第一个defer, r=n+1,
二、接着执行第二个defer，f()会报错
三、运行第三个defer，运行到recover，程序继续执行
*/
func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()

	defer f()
	f = func() {
		r += 2
	}
	return n + 1
}

func main() {
	fmt.Println(f(3))
}
