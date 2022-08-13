package main

import (
	"fmt"
	"strings"
)

/*
下面这段代码输出什么？请简要说明。
答案：结果为“”
解析：strings.TrimRight,会把第二个参数的所有组合都去掉，比如{"BA", "AB", "A", "B"}，遇到不匹配的才会停止
*/
func main() {
	fmt.Println(strings.TrimRight("ABBA", "BA"))
}
