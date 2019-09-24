// 7_2 Panic.go
package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaaaaaaaaa")
}

func testb(x int) {
	var a [10]int
	a[x] = 111 //当x为20时候，导致数组越界，产生一个panic，导致程序崩溃
}

func testb_delete() {
	//fmt.Println("bbbbbbbbbbbbbbbbbbbb")
	//显式调用panic函数，导致程序中断
	panic("this is a panic test")
}

func testc() {
	fmt.Println("cccccccccccccccccc")
}

func main() {
	testa()
	testb(20)
	testc()
}
