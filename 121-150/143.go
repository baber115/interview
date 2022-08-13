package main

/*
下面哪一行代码会 panic，请说明原因？
答案：	_ = s.bar
解析：因为 s.bar 将被展开为 (*s.T).bar，而 s.T 是个空指针，解引用会 panic。
*/

type T struct{}

func (*T) foo() {
}

func (T) bar() {
}

type S struct {
	*T
}

func main() {
	s := S{}
	_ = s.foo
	s.foo()
	_ = s.bar
}
