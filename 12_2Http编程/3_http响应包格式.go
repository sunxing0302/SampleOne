// 3_http响应包格式.go
package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("dial err =", err)
		return
	}

	defer conn.Close()

	requestBuf := "GET /test HTTP/1.1\r\nAccept: image/gif, image/jpeg, image/pjpeg, application/x-ms-application, application/xaml+xml, application/x-ms-xbap, */*\r\nAccept-Language: zh-Hans-CN,zh-Hans;q=0.8,en-US;q=0.5,en;q=0.3\r\nUser-Agent: Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)\r\nAccept-Encoding: gzip, deflate\r\nHost: 127.0.0.1:8000\r\nConnection: Keep-Alive\r\n\r\n"

	//先发请求包，服务器才会回响应包
	conn.Write([]byte(requestBuf))

	//接受服务器回复的响应包
	buf := make([]byte, 1024*4)
	n, err1 := conn.Read(buf)
	if n == 0 {
		fmt.Println("Read err =", err1)
		return
	}

	//打印响应报文
	fmt.Printf("#%v#", string(buf[:n]))
	//响应报文失败：
	/*
		HTTP/1.1 404 Not Found
		Content-Type: text/plain; charset=utf-8
		X-Content-Type-Options: nosniff
		Date: Sat, 04 May 2019 10:52:44 GMT
		Content-Length: 19

		404 page not found
	*/

	//响应报文 成功
	/*
		HTTP/1.1 200 OK
		Date: Sat, 04 May 2019 10:55:37 GMT
		Content-Length: 22
		Content-Type: text/plain; charset=utf-8

		this is a test server
	*/
}
