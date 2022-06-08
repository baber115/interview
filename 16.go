package main

import "fmt"

// 关于init，以下说法正确的是
/**
A.一个包里可以包含多个init方法
B.程序编译时，会先执行依赖包的init方法，再执行main的init方法
C.main包中，不能有init方法
D.init方法可以被其他方法调用

答案
AB
1.init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
2.一个包可以有多个init
3.同一个包中有多个init函数，按加载顺序执行，不同包的init函数根据包导入的依赖关系决定
4.init函数在代码中不能被显示调用、不能被引用（赋值给函数变量），否则出现编译失败；
5.一个包被多次调用，A import B,C import B, D import A,B被多次调用，但B包只会初始化一次
6.import包的时候不能死循环，如：A import B,B import A，会出现编译错误
*/

func init() {
	fmt.Println(1)
}

func init() {
	fmt.Println(2)
}

func main() {
	fmt.Println(3)
}
