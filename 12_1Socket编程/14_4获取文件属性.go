// 14_4获取文件属性.go
package main

import (
	"fmt"
	"os"
)

func main() {

	list := os.Args
	if len(list) != 2 {
		fmt.Println("useage xxx.exe file")
		return
	}
	fileName := list[1]
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Printf("Name : %s, Size : %d", fileInfo.Name(), fileInfo.Size())

}
