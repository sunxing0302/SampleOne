//重置定时器
package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)
	timer.Reset(time.Second) //重新设置为1s
	<-timer.C
	fmt.Println("时间到")
}
