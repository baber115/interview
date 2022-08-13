package main

import "fmt"

/*
下面代码输出什么？
答案：[1 2 3 4 5]
解析：a在for range过程上中加了两个元素，但是for range进使用的是a的拷贝参与循环，所以还是循环5次，并且获取a的底层数据
*/
func main() {
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)

	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}

		r = append(r, v)
	}

	fmt.Println(r)
}
