package main

import "fmt"

/**
下面输出的结果是否正确
输出：
{A} {B} {C}
&{A} &{B} &{C}

答案：错误
s2输出的应该是&{C} &{C} &{C}

解析：line:29，&value是指针，保留最后一次赋值的值
*/

type Foo struct {
	bar string
}

func main() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		s2[i] = &value
	}
	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2])
}
