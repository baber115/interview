package main

import "fmt"

// 以下代码能否通过编译

/**
不能通过,invalid operation: s3 == s4
考点结构体的比较
1、结构体只能比较是否相等，不能比较大小
2、相同类型的结构体才能比较，结体的属性值和属性顺序，都会影响判断的结果
3、结构体的所有成员可以互相比较，结构体才能比较是否相等，按顺序逐个比较，如果全相等，则结构体相等

常见的能比较的属性
bool、数值类型、字符、指针、数组

不能比较的属性
slice、map、函数

*/
func main() {
	s1 := struct {
		a int
		b string
	}{a: 1, b: "1"}
	s2 := struct {
		a int
		b string
	}{a: 1, b: "1"}

	if s1 == s2 {
		fmt.Println("s1 == s2")
	}

	s3 := struct {
		age int
		m   map[string]string
	}{age: 1, m: map[string]string{"a": "a"}}
	s4 := struct {
		age int
		m   map[string]string
	}{age: 2, m: map[string]string{"b": "b"}}
	if s3 == s4 {
		fmt.Println("s3==s4")
	}
}
