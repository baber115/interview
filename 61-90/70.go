package main

/**
以下代码输出什么
答案：
r=[1,2,3,4,5]
a=[1,12,13,4,5]
解析：
for range 是拷贝一份出来，不影响原来值
如果要修改原变量的值，可以用指针
for i, v := range &a {
	if i == 0 {
		a[1] = 12
		a[2] = 13
	}
	r[i] = v
}
*/
import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}
