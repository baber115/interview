package main

import "fmt"

// 以下会输出什么
/**
编译失败
只有接口类型才能switch 判断类型，i.(tyoe)，其中i是接口，type是固定的关键字，只有接口类型才可以使用类型选择
*/
func GetValue() int {
	return 1
}

func main() {
	i := GetValue()
	switch i.(type) {
	case int:
		fmt.Println("int")
	default:
		fmt.Println("default")
	}
}
