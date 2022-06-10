package main

/**
关于字符串连接，下面正确的是
A. str := 'abc' + '123'
B. str := "abc" + "123"
C. str := '123' + "abc"
D. fmt.Sprintf("abc%d", 123)

答案：
B,D
GO语言字符串用双引号表示，单引号表示字符
除了D选项的拼接方式外，还有strings.join(),buffer.writeString()等
*/
