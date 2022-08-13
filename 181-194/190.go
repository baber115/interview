package main

import "fmt"

/*
下面的代码有什么问题，请说明。
解析：var in printer = data{"two"}有问题
data没有实现print方法，&data实现了print方法,
修改后：
var in printer = &data{"two"}
*/
type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func main() {
	d1 := data{"one"}
	d1.print()

	var in printer = data{"two"}
	in.print()
}
