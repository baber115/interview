package main

import "fmt"

/*
下列代码输出什么
答案：
r =  [1 12 13 4 5]
a =  [1 12 13 4 5]
解析：
切片里有指针，在for range的时候发生拷贝，会拷贝原切片的指针，所以修改时会直接修改原数组上，通过v拿数据时拿到的就是修改后的值
*/
func main() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		fmt.Println(v)
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		fmt.Println(v)
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}
