package main

/*
下面代码能编译通过吗？可以的话，输出什么？
答案：A:true, B:false
解析：go代码断行规则，
A处是switch alwaysFalse();{，相当于switch alwaysFalse(); true{
*/
func alwaysFalse() bool {
	return false
}

func main() {
	// A
	switch alwaysFalse(); {
	case true:
		println(true)
	case false:
		println(false)
	}

	// B
	switch alwaysFalse() {
	case true:
		println(true)
	case false:
		println(false)
	}
}
