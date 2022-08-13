package main

/*
下面的代码有几处问题？请详细说明。
type T struct {
    n int
}

func (t *T) Set(n int) {
    t.n = n
}

func getT() T {
    return T{}
}

func main() {
    getT().Set(1)
}

答案：有两处问题
解析：
1、getT()返回的T不是指针
2、getT().Set(1)中，getT()是不可寻址的，不能调用方法
*/
