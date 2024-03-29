// TCP服务器.go
package main

import (
	"fmt"
	"net"
	"strings"
)

//处理用户请求
func HandleConn(conn net.Conn) {
	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println("addr connect sucessful")

	buf := make([]byte, 2048)

	for {
		//读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err =", err)
			return
		}

		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))

		//if string(buf[:n-1]) == "exit" { //nc 测试
		if string(buf[:n-2]) == "exit" { //自己写的客户端测试,输入时多了2个字符，"\r\n"
			fmt.Println("exit")
			return
		}
		//把数据转换为大写，再给用户发送
		str := strings.ToUpper(string(buf[:n]))
		conn.Write([]byte("server response : " + str))
	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	defer listener.Close()

	//接受多个用户
	for {
		conn, err := listener.Accept()
		for err != nil {
			fmt.Println("err =", err)
			return
		}

		//处理用户请求,新建一个协程
		go HandleConn(conn)

	}

}
