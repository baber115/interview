package main

import "fmt"

/**
下面的代码有几处语法问题，各是什么？

答案：
字符串类型不能赋nil值
nil不能用来做判断
*/
package main
import (
"fmt"
)
func main() {
	var x string = ""
	if x == nil {
		x = "default"
	}
	fmt.Println(x)
}
