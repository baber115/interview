package main

/*
函数执行时，如果由于 panic 导致了异常，则延迟函数不会执行。这一说法是否正确？
A. true
B. false
答：B
解析：由 panic 引发异常以后，程序停止执行，然后调用延迟函数（defer），就像程序正常退出一样。
*/
