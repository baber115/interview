package main

/*
请指出下面代码的错误？
答案：变量未使用
 */
 */
import "fmt"

package main

var gvar int

func main() {
	var one int
	two := 2
	var three int
	three = 3

	func(unused string) {
		fmt.Println("Unused arg. No compile error")
	}("what?")
}