package main

/*
下面的代码有什么问题？
type T struct {
    n int
}

func getT() T {
    return T{}
}

func main() {
    getT().n = 1
}
答案：错误
解析：getT()方法返回的无法寻址，不能赋值
*/
