package main

/*
下面的代码能否正确输出？
func main() {
    var fn1 = func() {}
    var fn2 = func() {}

    if fn1 != fn2 {
        println("fn1 not equal fn2")
    }
}

答案：编译错误
解析：函数只能与nil比较
*/
