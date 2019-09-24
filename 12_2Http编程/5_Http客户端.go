// 5_Http客户端.go
package main

import (
	"fmt"
	"net/http"
)

func main() {

	//resp, err := http.Get("http://www.baidu.com")
	resp, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		fmt.Println("http.Get err =", err)
		return

	}

	defer resp.Body.Close()

	/*
	   resp.Status =  200 OK
	   resp.StatusCode =  200
	   resp.Header =  map[Content-Length:[8] Content-Type:[text/plain; charset=utf-8] D
	   ate:[Sat, 04 May 2019 12:06:03 GMT]]
	   resp.Body =  &{0xc000052400 {0 0} false <nil> 0x60dd50 0x60dcd0}
	   read err =  EOF
	   tmp =  hello go
	*/
	fmt.Println("resp.Status = ", resp.Status)
	fmt.Println("resp.StatusCode = ", resp.StatusCode)
	fmt.Println("resp.Header = ", resp.Header)
	//fmt.Println("resp.Body = ", resp.Body)

	buf := make([]byte, 4*1024)
	var tmp string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("read err = ", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("resp.Body = ", tmp)

}
