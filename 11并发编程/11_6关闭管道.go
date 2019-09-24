package main

import (
	"fmt"
)

func main() {
	//创建一个无缓存的channel
	ch := make(chan int)

	//新建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往管道写数据
		}
		//不需要再写数据时，关闭channel
		close(ch)
	}()

	for {
		//如果ok为true，说明管道没有关闭
		num, ok := <-ch
		if ok == true {
			fmt.Println("num=", num)
		} else {
			fmt.Println("管道已经关闭")
			break
		}

	}
}
