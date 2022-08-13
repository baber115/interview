package main

/*
下面代码输出什么？
答案：{5}
解析：
这道题容易忽视的点是，String() 是指针方法，而不是值方法，所以使用 Println() 输出时不会调用到 String() 方法。
*fmt.Println会在内部调用Error()和String()方法,而且Error()方法有优先权
*/
import "fmt"

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

func (o *Orange) String() string {
	return fmt.Sprintf("%#v", o.Quantity)
}

func main() {
	var orange Orange
	//orange := &Orange{} 这样写会调用string方法
	orange.Increase(10)
	orange.Decrease(5)
	fmt.Println(orange)
}
