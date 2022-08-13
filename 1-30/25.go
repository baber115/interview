package main

import "fmt"

/**
关于 cap() 函数的适用类型，下面说法正确的是()
A. array
B. slice
C. map
D. channel

答案：ABD
cap不适用于map
*/
func main() {
	m := make(map[int]int, 10)
	fmt.Println(cap(m))
}
