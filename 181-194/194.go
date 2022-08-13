package main

import "fmt"

/*
下面代码输出什么？
答案：[]
解析：copy返回的是int值，是一个size，size=min(len(dst),len(src)),这个size决定要拷贝几个元素
由于声明了dst但是没有初始化，所以len(dst)=0，size=0，因此没有copy任何元素
*/
func main() {
	var src, dst []int
	src = []int{1, 2, 3}
	copy(dst, src)
	fmt.Println(dst)
}
