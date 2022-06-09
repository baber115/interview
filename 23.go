package main

import "fmt"

// 以下代码输出什么
/**
答案:4

[i:j]表示分割数组/切片
	1.从索引i到索引j的值,但不包括j
	2.i省略,默认从0开始.j省略,默认数组/切片的长度
截取操作还可以跟第3个值,用来限制切片的容量,但不能超过原切片的容量大小.截取获得的切片的长度和容量分别是：j-i、k-i。
*/
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[2:4:4]
	fmt.Println(t)
	fmt.Printf("len:%d\n", len(t))
	fmt.Printf("cap:%d\n", cap(t))
}
