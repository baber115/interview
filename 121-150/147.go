package main

/*
下面哪一行代码会panic
func main() {
    nil := 123
    fmt.Println(nil)
    var _ map[string]int = nil
}
答案：var _ map[string]int = nil
解析：在前面nil关键字被赋值成int类型，int类型不能赋值给map
*/
