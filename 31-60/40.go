package main

import "fmt"

/**
切片 a、b、c 的长度和容量分别是多少？
答案：
cap(a)=3
len(a)=0
cap(b)=3
len(b)=2
cap(c)=2
len(c)=1

解析：

截取操作有带 2 个或者 3 个参数，形如：[i:j] 和 [i:j:k]，假设截取对象的底层数组长度为 l。在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的长度，截取得到的切片长度和容量计算方法是 j-i、l-i。
操作符 [i:j:k]，k 主要是用来限制切片的容量，但是不能大于数组的长度 l，截取得到的切片长度和容量计算方法是 j-i、k-i。
*/
func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]

	fmt.Printf("cap(a)=%d,", cap(a))
	fmt.Printf("len(a)=%d\n", len(a))
	fmt.Printf("cap(b)=%d,", cap(b))
	fmt.Printf("len(b)=%d\n", len(b))
	fmt.Printf("cap(c)=%d,", cap(c))
	fmt.Printf("len(c)=%d\n", len(c))
}
