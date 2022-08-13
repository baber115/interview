package main

/*
下面代码能编译通过吗？
答案：编译失败，不能使用短声明设置结构体字段值
import "fmt"

type info struct {
	result int
}

func work() (int,error) {
	return 13,nil
}

func main() {
	var data info

	data.result, err := work()
	fmt.Printf("info: %+v\n",data)
}
*/
