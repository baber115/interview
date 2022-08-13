package main

/*
下面的代码有什么问题？
func main() {
    data := []int{1,2,3}
    i := 0
    ++i
    fmt.Println(data[i++])
}
解析：
1、Go语言中不支持++i和--i的操作
2、自增自减是独立的语句，不是表达式
*/
