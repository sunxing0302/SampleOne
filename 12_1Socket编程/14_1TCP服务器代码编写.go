// 14_1TCP服务器代码编写.go
package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	defer listener.Close()

	//阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err:", err)
			continue
		}

		//接受用户的请求
		buf := make([]byte, 1024) //1024大小的缓冲区
		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("err1:", err1)
			continue
		}

		fmt.Println("buf=", string(buf[:n]))

	}

}
