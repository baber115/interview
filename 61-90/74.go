package main

import "fmt"

/*
下面代码里的 counter 的输出值？
A. 2
B. 3
C. 2 或 3

答案：C

解析：for range map是无序的，如果第一次循环到A，则输出3，否则输出2
*/
func main() {

	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}
