// 11_3通过channel实现同步和数据交互.go
package main

import (
	"fmt"
	"time"
)

func main() {
	//创建channel
	ch := make(chan string)

	defer fmt.Println("主协程也结束")

	go func() {
		defer fmt.Println("子协程调用完毕")
		for i := 0; i < 2; i++ {
			fmt.Println("子协程 i = ", i)
			time.Sleep(time.Second)
		}

		ch <- "我是子协程，我工作完毕"

	}()

	str := <-ch //没有数据前，阻塞
	fmt.Println("str = ", str)

}
