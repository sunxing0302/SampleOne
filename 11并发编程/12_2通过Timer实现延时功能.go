//休眠2秒钟的3种方法
package main

import (
	"fmt"
	"time"
)

func main() {
	<-time.After(2 * time.Second) //定时2s，阻塞2s, 2s后产生一个事件，往channel写内容
	fmt.Println("时间到")
}

func main2() {
	time.Sleep(2 * time.Second)
	fmt.Println("时间到")
}

//延时5s后打印一句话
func main1() {
	fmt.Println(time.Now())
	timer := time.NewTimer(5 * time.Second)
	<-timer.C

	fmt.Println("时间到")
	//fmt.Println(timer)
	fmt.Println(time.Now())
}
