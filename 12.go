package main

import "fmt"

/**
[7 8 9]
[7 8 9]
[1 8 9]
append会重新分配内存空间，所以ap方法里的a，和外面的a，不是同一片内存空间，不会修改外面a的值
*/
func main() {
	a := []int{7, 8, 9}
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)
}

func ap(a []int) {
	a = append(a, 10)
}

func app(a []int) {
	a[0] = 1
}
