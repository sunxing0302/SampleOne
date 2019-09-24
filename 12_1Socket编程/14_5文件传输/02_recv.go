// 02_recv.go
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//接受文件内容
func RecvFile(fileName string, conn net.Conn) {
	//新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	buf := make([]byte, 1024*5)

	//接受多少，写多少，一点不差
	for {
		n, err := conn.Read(buf) //接受对象发过来的文件内容
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接受完毕")
			} else {
				fmt.Println("conn.Read err =", err)
			}
			return
		}

		if n == 0 {
			fmt.Println("n ==0 文件接受完毕")
		}
		f.Write(buf[:n]) //往文件写入内容
	}

}

func main() {
	//监听
	listenner, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	defer listenner.Close()

	//阻塞等待用户连接
	conn, err1 := listenner.Accept()
	if err1 != nil {
		fmt.Println("listenner.Accept err = ", err1)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	var n int
	n, err = conn.Read(buf) //读取对方发送的文件名
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}

	fileName := string(buf[:n])
	//回复“OK”
	conn.Write([]byte("OK"))

	//接受文件内容
	RecvFile(fileName, conn)

}
