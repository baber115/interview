package main

import "fmt"

/*
下面代码输出正确的是？
A. s: [Z,B,C]
B. s: [A,Z,C]

答案：A
解析：等同于 i, s[0] = 2, "Z"
*/
func main() {
	i := 1
	s := []string{"A", "B", "C"}
	i, s[i-1] = 2, "Z"
	fmt.Printf("s: %v \n", s)
}
